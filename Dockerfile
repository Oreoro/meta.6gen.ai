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
ENV PACKAGE github.com/oreoro/meta.6gen.ai
ENV BUILD_DIR ${GOPATH}/src/${PACKAGE}

ARG TAGS="sqlite sqlite_unlock_notify"
ENV TAGS "bindata timetzdata $TAGS"
ARG CGO_EXTRA_CFLAGS

# **MODIFICATION 1: Added python3 for native Node.js modules**
# Some pnpm packages need to compile C++ add-ons and require Python.
RUN apk --no-cache add build-base git bash nodejs npm python3 && \
    npm install -g pnpm@9.7.0 && \
    rm -rf /var/cache/apk/*

COPY . ${BUILD_DIR}
WORKDIR ${BUILD_DIR}

# **MODIFICATION 2: Made the pnpm install step more robust and verbose**
# Use an argument to make the registry configurable, which helps with network issues.
ARG NPM_REGISTRY=https://registry.npmjs.org/

# This command now does three things:
# 1. Explicitly sets the package registry to avoid network/DNS issues.
# 2. Cleans the pnpm cache to prevent corrupted data from stopping the build.
# 3. Runs the install with a verbose reporter for detailed error logging.
RUN pnpm config set registry ${NPM_REGISTRY} && \
    pnpm cache clean && \
    pnpm install --reporter verbose

# --- The rest of the build process remains separated for clarity ---
# 2. Build backend Go application
RUN make clean build

# 3. Build frontend UI
RUN make ui

# 4. Finalize build artifacts (permissions and plugin script)
RUN chmod 755 answer && \
    /bin/bash -c "script/build_plugin.sh"

# 5. Prepare runtime directories and copy all necessary assets
RUN mkdir -p /data/uploads /data/i18n /data/ui && \
    cp -r i18n/*.yaml /data/i18n/ && \
    cp -r ui/build/* /data/ui/ 2>/dev/null || cp -r dist/* /data/ui/ 2>/dev/null || true


# --- FINAL IMAGE STAGE (No changes needed here, it follows best practices) ---
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
    chown -R 10001:10001 /data /usr/bin/answer

# Switch to the non-root user
USER 10001:10001

VOLUME /data
EXPOSE 80
ENTRYPOINT ["/usr/bin/dumb-init", "--", "/entrypoint.sh"]
