# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements. See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.

# --- BUILDER STAGE ---
FROM golang:1.23-alpine AS golang-builder
LABEL maintainer="linkinstar@apache.org"

ARG GOPROXY
ENV GOPATH /go
ENV GOROOT /usr/local/go
# Using the package name from your original file
ENV PACKAGE github.com/oreoro/meta.6gen.ai
ENV BUILD_DIR ${GOPATH}/src/${PACKAGE}

ARG TAGS="sqlite sqlite_unlock_notify"
ENV TAGS "bindata timetzdata $TAGS"
ARG CGO_EXTRA_CFLAGS

# Install all build dependencies in a single layer for efficiency
RUN apk --no-cache add build-base git bash nodejs npm && \
    npm install -g pnpm@9.7.0 && \
    rm -rf /var/cache/apk/*

COPY . ${BUILD_DIR}
WORKDIR ${BUILD_DIR}

# --- MODIFIED BUILD PROCESS ---
# Each step is now a separate, debuggable layer.
# This structure is inspired by your working Dockerfile.

# 1. Install frontend dependencies
RUN pnpm install

# 2. Build backend Go application
RUN make clean build

# 3. Build frontend UI
RUN make ui

# 4. Finalize build artifacts (permissions and plugin script)
RUN chmod 755 answer && \
    /bin/bash -c "script/build_plugin.sh"

# 5. Prepare runtime directories and copy all necessary assets
# This includes the UI assets built by 'make ui'
RUN mkdir -p /data/uploads /data/i18n /data/ui && \
    cp -r i18n/*.yaml /data/i18n/ && \
    # This robust copy handles different possible UI output directories
    cp -r ui/build/* /data/ui/ 2>/dev/null || cp -r dist/* /data/ui/ 2>/dev/null || true


# --- FINAL IMAGE STAGE ---
# This stage uses the best practices from your original file (non-root user, dumb-init)
FROM alpine
LABEL maintainer="linkinstar@apache.org"

ARG TIMEZONE
ENV TIMEZONE=${TIMEZONE:-"Asia/Shanghai"}

# Security hardening and package installation
RUN apk update && apk --no-cache add \
        bash ca-certificates curl dumb-init \
        gettext openssh sqlite gnupg tzdata \
    && ln -sf /usr/share/zoneinfo/${TIMEZONE} /etc/localtime \
    && echo "${TIMEZONE}" > /etc/timezone \
    && sed -i -e 's/^root::/root:!:/' /etc/shadow \
    && rm -rf /var/cache/apk/*

# Create non-root user for better security
RUN addgroup -g 10001 -S appgroup && \
    adduser -u 10001 -S appuser -G appgroup

# Copy application binary and data from the builder stage
COPY --from=golang-builder ${BUILD_DIR}/answer /usr/bin/answer
COPY --from=golang-builder /data /data
COPY /script/entrypoint.sh /entrypoint.sh

# Set proper ownership and permissions for the non-root user
RUN chmod 755 /entrypoint.sh && \
    # The 'answer' binary needs to be owned by the user who will run it
    chown -R 10001:10001 /data /usr/bin/answer

# Switch to the non-root user
USER 10001:10001

VOLUME /data
EXPOSE 80
# Use dumb-init to properly handle signals, which is a best practice
ENTRYPOINT ["/usr/bin/dumb-init", "--", "/entrypoint.sh"]
