FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.20-alpine as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH
ARG VERSION="dev"

RUN apk update && apk add -U --no-cache ca-certificates

WORKDIR /app/
ADD go.mod go.sum ./
RUN go mod download
ADD main.go .
ADD pkg/ ./pkg
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s -X 'main.Version=${VERSION}'" -o tailscale-exporter main.go

FROM --platform=${TARGETPLATFORM:-linux/amd64} scratch
WORKDIR /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/tailscale-exporter /app/tailscale-exporter
ENTRYPOINT ["/app/tailscale-exporter"]
