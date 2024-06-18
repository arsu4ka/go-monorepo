# Build stage
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY ./internal/app2 ./internal/app2
COPY ./cmd/app2 ./cmd/app2
COPY ./pkg ./pkg

RUN go mod tidy
RUN go mod download

RUN go build -o /app/bin/app2 ./cmd/app2/main.go

# Run stage
FROM debian:latest

WORKDIR /app

COPY --from=builder /app/bin/app2 /app/bin/app2

RUN apt-get update -y && apt-get install ca-certificates -y

EXPOSE 8082
CMD ["/app/bin/app2"]
