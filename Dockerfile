# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

# -------------------- STAGE 1: BUILD THE APPLICATION (golang-builder) --------------------
FROM golang:1.23-alpine AS golang-builder
LABEL maintainer="linkinstar@apache.org"

ARG GOPROXY
ENV GOPATH /go
ENV GOROOT /usr/local/go

# 1. 🚨 CRITICAL FIX for build-args: Define PACKAGE as an ARG with your path
ARG PACKAGE=github.com/Oreoro/meta.6gen.ai
ENV PACKAGE ${PACKAGE}
ENV BUILD_DIR ${GOPATH}/src/${PACKAGE}
ENV ANSWER_MODULE ${BUILD_DIR}

ARG TAGS="sqlite sqlite_unlock_notify"
ENV TAGS "bindata timetzdata $TAGS"
ARG CGO_EXTRA_CFLAGS

COPY . ${BUILD_DIR}
WORKDIR ${BUILD_DIR}
# Install necessary dependencies and build both Go (answer binary) and Node (UI/front-end) assets
RUN apk --no-cache add build-base git bash nodejs npm && npm install -g pnpm@9.7.0 \
    && make clean build

# Finalize binary and static resources paths in the build stage
RUN chmod 755 answer
# Build and prepare plugins
RUN ["/bin/bash","-c","script/build_plugin.sh"]
RUN cp answer /usr/bin/answer

# Create and copy config/data files to the /data mount point
RUN mkdir -p /data/uploads && chmod 777 /data/uploads \
    && mkdir -p /data/i18n && cp -r i18n/*.yaml /data/i18n

# -------------------- STAGE 2: CREATE THE FINAL RUNTIME IMAGE --------------------
FROM alpine
LABEL maintainer="linkinstar@apache.org"

ARG TIMEZONE
ENV TIMEZONE=${TIMEZONE:-"Asia/Shanghai"}

# Install minimal runtime dependencies
RUN apk update \
    && apk --no-cache add \
        bash \
        ca-certificates \
        curl \
        dumb-init \
        gettext \
        openssh \
        sqlite \
        gnupg \
        tzdata \
    && ln -sf /usr/share/zoneinfo/${TIMEZONE} /etc/localtime \
    && echo "${TIMEZONE}" > /etc/timezone

# 1. Copy the compiled Go binary
COPY --from=golang-builder /usr/bin/answer /usr/bin/answer

# 2. Copy the configuration/i18n data
COPY --from=golang-builder /data /data

# 3. 🚨 CRITICAL FIX for UI Assets: Copy the built UI assets (index.html, JS, CSS, etc.)
# This resolves the "open build/index.html: file does not exist" error.
COPY --from=golang-builder ${BUILD_DIR}/ui/build /build

# 4. Copy and set executable permissions on the entrypoint script
COPY /script/entrypoint.sh /entrypoint.sh
RUN chmod 755 /entrypoint.sh

VOLUME /data
EXPOSE 80
ENTRYPOINT ["/entrypoint.sh"]
