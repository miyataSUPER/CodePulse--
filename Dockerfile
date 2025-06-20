# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.21 AS builder
ARG VERSION
WORKDIR /src
COPY . .
RUN go build -o /work/codepulse ./cmd/main

# Final stage
FROM scratch
ARG VERSION
LABEL org.opencontainers.image.source="https://github.com/miyataSUPER/CodePulse--" \
      org.opencontainers.image.version="${VERSION}" \
      org.opencontainers.image.title="codepulse" \
      org.opencontainers.image.description="File type classifier"
RUN adduser --disabled-password --disabled-login --home /workdir nonroot \
    && mkdir -p /workdir
COPY --from=builder /work/codepulse /opt/codepulse/codepulse
COPY --from=golang:1.21 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
WORKDIR /workdir
USER nonroot
ENTRYPOINT ["/opt/codepulse/codepulse"]
