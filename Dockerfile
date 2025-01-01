# Etapa de build
FROM golang:1.23 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o climate ./cmd/server

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/climate .
COPY --from=builder /app/.env /app/.env
COPY secret_key.txt /app/secret_key.txt


RUN apk add --no-cache bash

ENV PATH="/app:${PATH}"

ENTRYPOINT ["/app/climate"]
