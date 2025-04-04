ARG GO_VERSION=1.23.4
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /workspace
COPY . .
RUN go install github.com/go-delve/delve/cmd/dlv@latest
