FROM golang:1.24 AS builder

WORKDIR /go/src/github.com/unkeyed/unkey/go
# Copy everything first because go.mod has replace directives pointing to ./deploy/pkg/ modules
# This ensures all local dependencies are available before go mod download
COPY . .
RUN go mod download
ARG VERSION
ENV CGO_ENABLED=0
RUN go build -o bin/unkey -ldflags="-X 'github.com/unkeyed/unkey/go/pkg/version.Version=${VERSION}'" ./main.go

FROM gcr.io/distroless/static-debian12
COPY --from=builder /go/src/github.com/unkeyed/unkey/go/bin/unkey /

LABEL org.opencontainers.image.source=https://github.com/unkeyed/unkey/go
LABEL org.opencontainers.image.description="Unkey API"

ENTRYPOINT  ["/unkey"]
