# 💸 FinTech Solutions Backend Challenge

## 🚀 Project Overview

This repository contains a backend solution for the technical challenge proposed by **FinTech Solutions Inc.**. The goal is to build two API services in **Go**:

1. **Migration Service** (`POST /migrate`)  
   Uploads and processes historical transaction data from CSV files into a PostgreSQL database.

2. **Balance Service** (`GET /users/{user_id}/balance`)  
   Retrieves a user's balance summary, with optional date filtering.

The system is designed to be robust, scalable, and easy to deploy locally using Docker.

---

## 🏗️ Architecture and Design Patterns

The project follows a **layered architecture** to promote separation of concerns and modularity:

- **Presentation Layer (Handlers)**: Handles HTTP requests and delegates logic to services.
- **Service Layer (Business Logic)**: Encapsulates core rules like CSV parsing and balance calculations.
- **Repository Layer (Data Access)**: Abstracts database operations using the **Repository Pattern**.

### ✅ SOLID Principles

- **SRP**: Each component has a single responsibility.
- **DIP**: Services depend on interfaces, not concrete implementations.

### 🧩 Additional Patterns

- **Repository Pattern**: Decouples persistence from business logic.
- **Service Layer Pattern**: Keeps transport logic (HTTP) separate from core rules.

---

## 🛠️ Technologies Used

- **Language**: Go (Golang)
- **Database**: PostgreSQL
- **Containers**: Docker & Docker Compose
- **Libraries**:
  - `go-chi/chi` — HTTP routing
  - `rs/zerolog` — Structured logging
  - `sethvargo/go-envconfig` — Environment config
  - `pgx` & `database/sql` — PostgreSQL drivers
  - `encoding/csv` — CSV parsing
  - `swaggo/swag` — Swagger documentation

---

## 🚀 Getting Started

### ▶️ Prerequisites

- Docker & Docker Compose installed

### ▶️ Run with Docker

```bash
git clone https://github.com/Sirpyerre/fintech-backend.git
cd fintech-backend
docker-compose up --build
```

The API will be available at:  
`http://localhost:8080`

---

## 📖 API Documentation

### 🔄 Migration Service

**`POST /migrate`**

- Upload a CSV file with transactions.
- Header: `Content-Type: multipart/form-data`
- Field name: `file`
- Expected columns: `id,user_id,amount,datetime`

### 💰 Balance Service

**`GET /users/{user_id}/balance`**

- Returns balance, total credits, and total debits.

**Optional query params**:
- `from`: Start date (RFC3339)
- `to`: End date (RFC3339)

### ✅ Sample Response

```json
{
  "balance": 25.21,
  "total_debits": 10,
  "total_credits": 15
}
```

📎 Swagger UI: [`/swagger/index.html`](http://localhost:8080/swagger/index.html)

---

## 📊 Observability

- **Structured logging** with `zerolog`
---

## 🧪 Testing

Run unit tests for services:

```bash
go test ./internal/services/...
```

Includes coverage for:
- Migration logic
- Balance calculations
- Error handling

---

## 🧠 Future Improvements

- Add Redis caching for frequent balance queries
- Improve CSV validation and error reporting
- Add integration tests with real DB
- Implement graceful shutdown and health checks
- Extend Swagger with response examples and error codes

---

## 📄 License

MIT — free to use, modify, and share.