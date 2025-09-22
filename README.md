# FinTech Solutions Backend

Updated README to reflect latest changes (September 2025).

## Overview

This repository implements a small backend in Go that provides two main features:

- A Migration endpoint to upload historical transactions (CSV) into a PostgreSQL database.
- A Balance endpoint to query user balances (optionally filtered by date range).

The project is intended to be run locally using Docker / Docker Compose or directly via `go build`.

## Notable updates

- Application HTTP port changed to 8000 (API listens on :8000 and docker-compose maps host 8000 -> container 8000).
- Environment variables and defaults are loaded via `sethvargo/go-envconfig`.
- Swagger UI is exposed at `/swagger/`.
- Health endpoint at `/health` (used by Docker healthchecks).

## Technologies

- Go 1.24
- PostgreSQL (containerized)
- Docker & Docker Compose
- chi router, zerolog, pgx/database/sql

## Quickstart (Docker)

Prerequisites: Docker and Docker Compose.

1. Build and start services

```bash
docker-compose up --build
```

2. The API will be available at http://localhost:8000

3. Healthcheck: http://localhost:8000/health

4. Swagger UI: http://localhost:8000/swagger/index.html

## Environment variables (used in docker-compose)

- DATABASE_URL: postgres://fintech_user:secure_password_123@db:5432/fintech_db
- PORT: 8000 (default)
- LOG_LEVEL: debug
- ENV: development

When running locally without Docker, set `DATABASE_URL` and optionally `PORT`/`LOG_LEVEL`.

## Build and run locally

1. Install dependencies and build

```bash
## FinTech Backend

This repository contains a small Go backend that provides two main features:

- Uploading historical transactions (CSV) into a PostgreSQL database via a migration endpoint.
- Querying user balances via a balance endpoint with optional date range filtering.

The service runs on port 8000 by default and can be run with Docker Compose or built and run locally.

## What's in this README

- Setup (Docker and local)
- Configuration / Environment variables
- API endpoints and examples
- Database & migrations
- Testing

## Requirements

- Go 1.24
- Docker & Docker Compose (for containerized run)

## Quickstart — Docker (recommended)

1. Build and start services

```bash
docker-compose up --build
```

2. The API will be available at:

http://localhost:8000

3. Useful endpoints

- Health: http://localhost:8000/health
- Swagger UI: http://localhost:8000/swagger/index.html

The `docker-compose.yml` maps host port 8000 to container port 8000 and mounts `init.sql` into the Postgres container to initialize the schema on first boot.

## Run locally (without Docker)

1. Ensure you have a running PostgreSQL instance and set `DATABASE_URL` accordingly.

2. Download dependencies and build:

```bash
go mod download
go build -o fintech-backend ./cmd/api
```

2. Run (example):

```bash
DATABASE_URL=postgres://fintech_user:secure_password_123@localhost:5432/fintech_db \
PORT=8000 LOG_LEVEL=debug ./fintech-backend
```

## Configuration / Environment variables

The app uses `sethvargo/go-envconfig` to load environment variables. Key variables:

- DATABASE_URL (required) — Postgres connection string, e.g. postgres://user:pass@host:5432/dbname
- PORT — HTTP port (default: 8000)
- LOG_LEVEL — logging level (default: info)
- ENV — environment name (default: development)

Example `docker-compose` environment (used in the included `docker-compose.yml`):

```yaml
DATABASE_URL: postgres://fintech_user:secure_password_123@db:5432/fintech_db
PORT: 8000
LOG_LEVEL: debug
ENV: development
```

## API Reference

Base URL: http://localhost:8000

1) POST /migrate

- Description: Upload a CSV file containing historical transactions which will be parsed and inserted into the database.
- Request: multipart/form-data; form field name: `file`
- CSV columns expected (header): id,user_id,amount,datetime
- Example:

```bash
curl -v -F "file=@history.csv" http://localhost:8000/migrate
```

2) GET /users/{user_id}/balance

- Description: Return the current balance for a user. Supports optional `from` and `to` query parameters to restrict transactions considered.
- Query params (optional):
  - from (RFC3339 or YYYY-MM-DD)
  - to (RFC3339 or YYYY-MM-DD)
- Examples:

```bash
# Latest balance for user 1
curl http://localhost:8000/users/1/balance

# Balance for user 1 between dates
curl "http://localhost:8000/users/1/balance?from=2023-01-01&to=2023-12-31"
```

Response (example JSON):

```json
{
  "balance": 25.21,
  "total_debits": 10,
  "total_credits": 15
}
```

3) GET /health

- Returns 200 OK when the service is available.

4) Swagger UI

- Available at: /swagger/index.html

## Database & migrations

- `init.sql` contains the schema used during initial database startup (mounted into the Postgres container).
- For manual DB setup, run the SQL in `init.sql` against your Postgres instance before starting the service.

## Testing

- Unit tests for services exist in `internal/services` (e.g., `balance_service_test.go`, `migration_service_test.go`). Run them with:

```bash
go test ./...
```

There are service-level tests under `internal/services` (e.g., `balance_service_test.go`, `migration_service_test.go`).

## Development notes

- Config is loaded with `sethvargo/go-envconfig` in `internal/config/config.go`.
- Server entrypoint: `cmd/api/main.go` — registers the routes:
  - POST `/migrate`
  - GET `/users/{user_id}/balance`
  - GET `/health`
  - GET `/swagger/*`
- Dockerfile builds the binary with Go 1.24 and serves it from an alpine image on port 8000.

## Troubleshooting

- Database connection errors: verify `DATABASE_URL` and that Postgres is accepting connections.
- Port in use: ensure nothing else is bound to port 8000 or set `PORT` to another value.

## License

This project is provided as-is. See repository for license details.

