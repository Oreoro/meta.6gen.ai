# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements. See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.

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

# Install build dependencies in single layer
RUN apk --no-cache add build-base git bash nodejs npm && \
    npm install -g pnpm@9.7.0 && \
    rm -rf /var/cache/apk/*

COPY . ${BUILD_DIR}
WORKDIR ${BUILD_DIR}

# Build both backend and frontend
RUN pnpm install && \
    make clean build && \
    make ui && \
    chmod 755 answer && \
    /bin/bash -c "script/build_plugin.sh"

# Prepare runtime directories
RUN mkdir -p /data/uploads /data/i18n /data/ui && \
    cp -r i18n/*.yaml /data/i18n/ && \
    cp -r ui/build/* /data/ui/ 2>/dev/null || cp -r dist/* /data/ui/ 2>/dev/null || true


# Use specific Alpine version with digest for security
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

# Create non-root user
RUN addgroup -g 10001 -S appgroup && \
    adduser -u 10001 -S appuser -G appgroup

# Copy application files
COPY --from=golang-builder /go/src/github.com/oreoro/meta.6gen.ai/answer /usr/bin/answer
COPY --from=golang-builder /data /data
COPY /script/entrypoint.sh /entrypoint.sh

# Set proper permissions
RUN chmod 755 /entrypoint.sh && \
    chmod 755 /data/uploads && \
    chown -R 10001:10001 /data /usr/bin/answer

# Switch to non-root user
USER 10001:10001

VOLUME /data
EXPOSE 80
ENTRYPOINT ["/usr/bin/dumb-init", "--", "/entrypoint.sh"]
