# Build stage
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY ./internal/app1 ./internal/app1
COPY ./cmd/app1 ./cmd/app1
COPY ./pkg ./pkg

RUN go mod tidy
RUN go mod download

RUN go build -o /app/bin/app1 ./cmd/app1/main.go

# Run stage
FROM debian:latest

WORKDIR /app

COPY --from=builder /app/bin/app1 /app/bin/app1

RUN apt-get update -y && apt-get install ca-certificates -y
CMD ["/app/bin/app1"]
