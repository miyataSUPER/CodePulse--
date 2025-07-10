# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.20 AS builder
ARG VERSION
WORKDIR /src
COPY . .
RUN go build -o /work/codepulse ./cmd/main

# Final stage
FROM alpine:3.18 AS runtime
ARG VERSION
LABEL org.opencontainers.image.source="https://github.com/miyataSUPER/CodePulse--" \
      org.opencontainers.image.version="${VERSION}" \
      org.opencontainers.image.title="codepulse" \
      org.opencontainers.image.description="File type classifier"
RUN adduser -D -h /workdir nonroot \
    && mkdir -p /workdir
COPY --from=builder /work/codepulse /opt/codepulse/codepulse
COPY --from=golang:1.20 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
WORKDIR /workdir
USER nonroot
ENTRYPOINT ["/opt/codepulse/codepulse"]
