FROM golang:1.23.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go
FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/internal/config/db/migrations /app/internal/config/db/migrations

EXPOSE 8080
CMD ["./main"]
