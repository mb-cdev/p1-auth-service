FROM golang:1.21.0-bookworm AS builder
WORKDIR /src
COPY --from=src . .
RUN ["go", "get"]
RUN [ "go", "build", "-o", "authService", "main.go" ]

FROM debian:bookworm
WORKDIR /app
COPY --from=builder /src/authService .
ENTRYPOINT ["./authService"]