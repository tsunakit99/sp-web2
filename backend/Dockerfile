FROM golang:1.24.2 AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main ./cmd/main.go

# Run stage
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
