# Builder
FROM golang:1.24-alpine AS builder

WORKDIR /metrics

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go

# Runtime
FROM alpine:latest

WORKDIR /weekcase

COPY --from=builder /metrics/server .
COPY --from=builder /metrics/assets ./assets
COPY --from=builder /metrics/pages  ./pages