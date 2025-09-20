# FinTech Solutions Backend Challenge

## 🚀 Project Overview

This project is a solution to the backend development challenge proposed by FinTech Solutions Inc. Its main goal is to create two API services in **Go**:
1. A **Migration Service** (`/migrate`) to process and store historical transaction data from CSV files into a database.
2. A **Balance Service** (`/users/{user_id}/balance`) to retrieve a user's balance, including a date range filter.

The solution is designed to be robust, scalable, and easy to deploy in a local environment using Docker.

## 🏗️ Architecture and Design Patterns

The application is built on a **layered architecture** that promotes separation of concerns and modularity. The structure is divided into the following main layers:

* **Presentation Layer (Handlers)**: Receives HTTP requests, validates input data, and delegates business logic to the service layer.
* **Service Layer (Business Logic)**: Contains the core business rules, such as CSV processing and balance calculation.
* **Repository Layer (Data Access)**: Abstracts database interactions by implementing the **Repository Pattern**. This ensures the business logic remains independent of the underlying database technology.

---

### **SOLID Principles**

* **Single Responsibility Principle (SRP)**: Each component has only one reason to change (e.g., a handler only handles web requests, a service only handles business logic).
* **Dependency Inversion Principle (DIP)**: High-level modules (services) do not depend on low-level modules (database implementation); both depend on abstractions (interfaces).

### **Additional Design Patterns**

* **Repository Pattern**: Isolates the business layer from persistence details, making it easy to swap databases.
* **Service Layer Pattern**: Separates business logic from the transport (HTTP), resulting in cleaner and more reusable code.

## 🛠️ Technologies Used

* **Language**: **Go (Golang)**
* **Database**: **PostgreSQL**
* **Containers**: **Docker** and **Docker Compose** for easy deployment.
* **Go Frameworks and Libraries**:
    * **`go-chi/chi`**: A lightweight HTTP router for managing endpoints.
    * **`sethvargo/go-envconfig`**: For loading configuration from environment variables.
    * **`rs/zerolog`**: For high-performance, structured logging.
    * **`database/sql` & `pgx`**: Drivers for PostgreSQL interaction.
    * **`encoding/csv`**: Go's standard package for processing CSV files.

## 🚀 How to Get Started

Make sure you have **Docker** and **Docker Compose** installed on your system.

1.  Clone the repository:
    ```bash
    git clone <REPOSITORY_URL>
    cd <REPOSITORY_NAME>
    ```

2.  Start the containers:
    ```bash
    docker-compose up --build
    ```
    This command will build the Go application image and start both the API service and the PostgreSQL database. The application will be available at `http://localhost:8080`.

## 📖 API Documentation

The API exposes the following endpoints:

### **1. Migration Service**

* **`POST /migrate`**
    * **Description**: Processes a CSV file of transactions and stores them in the database.
    * **Header**: `Content-Type: multipart/form-data`
    * **Request Body**: A CSV file named `file` with the columns `id,user_id,amount,datetime`.

### **2. Balance Service**

* **`GET /users/{user_id}/balance`**
    * **Description**: Returns the total balance, total debits, and total credits for a specific user.
* **`GET /users/{user_id}/balance?from=...&to=...`**
    * **Description**: Returns the same balance, but filtered by a date range. Dates must be in ISO 8601 format (`YYYY-MM-DDThh:mm:ssZ`).

---

### **Response Format (for the balance service)**

```json
{
  "balance": 25.21,
  "total_debits": 10,
  "total_credits": 15
}