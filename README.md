# golang-monorepo

This repository contains two Go services (`app1` and `app2`) structured as a monorepo. The services communicate with each other using the Asynq package, which leverages Redis for task queuing. `app1` processes tasks sent from `app2`, and `app2` is a Telegram bot built using the telego and Fiber packages.

## Project Structure

```
.
├── cmd
│   ├── app1
│   │   └── main.go
│   └── app2
│       └── main.go
├── internal
│   ├── app1
│   │   ├── app.go
│   │   └── handlers.go
│   └── app2
│       ├── app.go
│       └── utils.go
├── pkg
│   ├── config
│   │   └── redis.go
│   └── tasks
│       └── user_joined.go
├── deploy
│   ├── app1
│   │   └── Dockerfile
│   └── app2
│       └── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
```

## Services

### app1

`app1` is responsible for processing tasks received from `app2`. It utilizes the Asynq package to handle these tasks in a robust and efficient manner.

### app2

`app2` is a Telegram bot built using the Telego package (a wrapper for the Telegram API) and Fiber (a web framework). The bot operates using webhooks to handle incoming updates.

## Communication

The services communicate through Redis using the Asynq package. `app2` queues tasks to be processed by `app1`.

## Setup

### Prerequisites

- Docker
- Docker Compose
- Go (version specified in `go.mod`)

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/arsu4ka/golang-monorepo.git
   cd golang-monorepo
   ```

2. Start the services using Docker Compose:
   ```
   docker-compose up --build
   ```

This will build and start both `app1` and `app2` services along with Redis.

### Running Locally

To run the applications locally without Docker, follow these steps:

1. Ensure Redis is running locally or adjust the configuration to point to your Redis instance.
2. Run `app1`:
   ```
   go run cmd/app1/main.go
   ```
3. Run `app2`:
   ```
   go run cmd/app2/main.go
   ```
   