# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements. See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.

# --- BUILDER STAGE ---
FROM golang:1.23-alpine AS golang-builder
LABEL maintainer="linkinstar@apache.org"

ARG GOPROXY
ARG PACKAGE=github.com/oreoro/meta.6gen.ai
ARG TAGS="sqlite sqlite_unlock_notify"
ARG CGO_EXTRA_CFLAGS
ARG NPM_REGISTRY=https://registry.npmjs.org/

ENV GOPATH=/go
ENV GOROOT=/usr/local/go
ENV PACKAGE=${PACKAGE}
ENV BUILD_DIR=/go/src/${PACKAGE}
ENV TAGS="bindata timetzdata ${TAGS}"

# Install build dependencies including make
RUN apk --no-cache add \
    build-base \
    git \
    bash \
    nodejs \
    npm \
    python3 \
    make && \
    rm -rf /var/cache/apk/*

# Set working directory first
RUN mkdir -p ${BUILD_DIR}
WORKDIR ${BUILD_DIR}

# Copy Go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Configure npm and install UI dependencies with sensible defaults and retries
ENV NODE_OPTIONS=--max-old-space-size=4096
WORKDIR ${BUILD_DIR}/ui
RUN npm config set registry "${NPM_REGISTRY:-https://registry.npmjs.org/}" && \
    npm config set fetch-retries 5 && \
    npm config set fetch-retry-mintimeout 20000 && \
    npm config set fetch-retry-maxtimeout 120000 && \
    npm config set legacy-peer-deps true && \
    npm config set ignore-scripts true && \
    npm cache clean --force

# Copy UI package files first for better dependency caching
COPY ui/package*.json ./
# Use npm install universally (skip prepare scripts) to avoid QEMU/pnpm issues
RUN npm install --no-audit --no-fund --legacy-peer-deps --loglevel=error || \
    (echo "npm install failed, retrying with forced legacy peer deps" && \
     npm install --no-audit --no-fund --legacy-peer-deps --force --loglevel=error)

# Return to repository root and copy all source code
WORKDIR ${BUILD_DIR}
COPY . .

# Build backend Go application
RUN make clean && make build

# Verify binary was created
RUN if [ ! -f "answer" ]; then \
        echo "ERROR: Binary 'answer' not found after build!" && \
        ls -la && \
        exit 1; \
    fi && \
    chmod 755 answer

# Build frontend UI (use npm directly instead of pnpm via make ui)
WORKDIR ${BUILD_DIR}/ui
RUN npm run build
WORKDIR ${BUILD_DIR}

# Prepare build artifacts (normalize binary name if needed)
RUN if [ -f "new_answer" ]; then mv new_answer answer; fi && \
    chmod 755 answer

# Build plugins if script exists (plugins are optional, so we handle failures gracefully)
# Note: build_plugin.sh uses 'set -e' and requires pnpm, so we need to handle failures carefully
# Skip plugin build if pnpm is not available (we're using npm-only build)
RUN if [ -f "script/build_plugin.sh" ] && [ -f "script/plugin_list" ] && [ -x "answer" ] && command -v pnpm >/dev/null 2>&1; then \
        echo "Building plugins..." && \
        (bash script/build_plugin.sh 2>&1 || echo "Plugin build failed, continuing without plugins") && \
        if [ -f "new_answer" ]; then \
            echo "Plugin build successful, replacing binary" && \
            rm -f answer && \
            mv new_answer answer && \
            chmod 755 answer; \
        else \
            echo "Plugin build did not produce new binary, continuing without plugins"; \
        fi; \
    else \
        if [ ! -f "script/build_plugin.sh" ] || [ ! -f "script/plugin_list" ]; then \
            echo "Skipping plugin build (script or plugin_list not found)"; \
        elif ! command -v pnpm >/dev/null 2>&1; then \
            echo "Skipping plugin build (pnpm not available, using npm-only build)"; \
        else \
            echo "Skipping plugin build (binary not executable)"; \
        fi; \
    fi

# Create entrypoint script if it doesn't exist
RUN if [ ! -f "script/entrypoint.sh" ]; then \
        mkdir -p script && \
        echo '#!/bin/sh' > script/entrypoint.sh && \
        echo 'exec /usr/bin/answer "$@"' >> script/entrypoint.sh && \
        chmod +x script/entrypoint.sh; \
    fi

# Create runtime directories and copy assets
RUN mkdir -p /data/uploads /data/i18n /data/ui && \
    if [ -d "i18n" ]; then cp -r i18n/*.yaml /data/i18n/ 2>/dev/null || true; fi && \
    if [ -d "ui/build" ]; then \
        cp -r ui/build/* /data/ui/; \
    elif [ -d "dist" ]; then \
        cp -r dist/* /data/ui/; \
    else \
        echo "Warning: No UI build artifacts found"; \
    fi

# --- FINAL IMAGE STAGE ---
FROM alpine:3.19
LABEL maintainer="linkinstar@apache.org"

ARG TIMEZONE
ENV TIMEZONE=${TIMEZONE:-Asia/Shanghai}
ARG PACKAGE=github.com/oreoro/meta.6gen.ai

# Install runtime dependencies and configure timezone
RUN apk update && apk --no-cache add \
        bash \
        ca-certificates \
        curl \
        dumb-init \
        gettext \
        openssh \
        sqlite \
        gnupg \
        tzdata && \
    ln -sf /usr/share/zoneinfo/${TIMEZONE} /etc/localtime && \
    echo "${TIMEZONE}" > /etc/timezone && \
    sed -i -e 's/^root::/root:!:/' /etc/shadow && \
    rm -rf /var/cache/apk/*

# Create non-root user for security
RUN addgroup -g 10001 -S appgroup && \
    adduser -u 10001 -S appuser -G appgroup

# Create required directories with proper permissions
RUN mkdir -p /data/uploads /data/i18n /data/ui && \
    chown -R 10001:10001 /data

# Explicitly re-declare PACKAGE in this stage if needed; usually not required unless reusing path variables.

# Copy application binary and data from builder
COPY --from=golang-builder --chown=10001:10001 /go/src/${PACKAGE}/answer /usr/bin/answer
COPY --from=golang-builder --chown=10001:10001 /data /data

# Copy entrypoint script (always exists after builder stage preparation)
COPY --from=golang-builder --chown=10001:10001 /go/src/${PACKAGE}/script/entrypoint.sh /entrypoint.sh
RUN chmod 755 /entrypoint.sh

# Copy verification script (for host-side verification, also included in image for reference)
COPY --from=golang-builder --chown=10001:10001 /go/src/${PACKAGE}/verify_custom_ui.sh /usr/local/bin/verify_custom_ui.sh
RUN chmod 755 /usr/local/bin/verify_custom_ui.sh

# Switch to non-root user
USER 10001:10001

VOLUME /data

# Use port 8080 instead of 80 for non-root user
EXPOSE 8080

ENTRYPOINT ["/usr/bin/dumb-init", "--", "/entrypoint.sh"]
