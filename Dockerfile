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

# Copy source code
COPY . .

# Configure npm and install UI dependencies with sensible defaults and retries
ENV NODE_OPTIONS=--max-old-space-size=4096
WORKDIR ${BUILD_DIR}/ui
RUN npm config set registry "${NPM_REGISTRY:-https://registry.npmjs.org/}" && \
    npm config set fetch-retries 5 && \
    npm config set fetch-retry-mintimeout 20000 && \
    npm config set fetch-retry-maxtimeout 120000 && \
    npm config set legacy-peer-deps true && \
    npm cache clean --force && \
    npm install -g pnpm@9

# Install using lockfile when available
COPY ui/package*.json ./
# Use npm install universally to avoid environments lacking `npm ci`
RUN npm install --no-audit --no-fund --legacy-peer-deps --loglevel=error || \
    (echo "npm install failed, retrying with forced legacy peer deps" && \
     npm install --no-audit --no-fund --legacy-peer-deps --force --loglevel=error)

# Copy UI sources after dependencies to leverage layer caching
COPY ui/ ./

# Return to repository root for Go build and make targets
WORKDIR ${BUILD_DIR}

# Build backend Go application
RUN make clean && make build

# Build frontend UI
RUN make ui

# Prepare build artifacts
RUN chmod 755 answer && \
    if [ -f "script/build_plugin.sh" ]; then \
        bash script/build_plugin.sh; \
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
COPY --from=golang-builder --chown=10001:10001 /go/src/github.com/oreoro/meta.6gen.ai/answer /usr/bin/answer
COPY --from=golang-builder --chown=10001:10001 /data /data

# Copy entrypoint script if it exists
COPY --chown=10001:10001 script/entrypoint.sh /entrypoint.sh
RUN chmod 755 /entrypoint.sh

# Switch to non-root user
USER 10001:10001

VOLUME /data

# Use port 8080 instead of 80 for non-root user
EXPOSE 8080

ENTRYPOINT ["/usr/bin/dumb-init", "--", "/entrypoint.sh"]
