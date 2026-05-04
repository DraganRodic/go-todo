# 📝 Go Todo API

A RESTful API built with Go for managing todos with user authentication.

---

## 🚀 Features

* User registration and login (JWT authentication)
* Protected routes using middleware
* Full CRUD operations for todos
* Soft delete support
* Request validation (email, password, required fields)
* Global error handling (consistent API responses)
* Pagination (`page`, `limit`)
* Filtering (`completed`)
* Sorting (`asc`, `desc`)
* Swagger API documentation

---

## 🛠️ Tech Stack

* **Go (Golang)**
* **Gin** – HTTP web framework
* **GORM** – ORM for database handling
* **MySQL** – relational database
* **JWT** – authentication

---

## ⚙️ Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/DraganRodic/go-todo.git
cd go-todo
```

---

### 2. Create `.env` file

```env
PORT=8080

DB_USER=root
DB_PASS=yourpassword
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=todo_db

JWT_SECRET=your_secret_key
```

---

### 3. Install dependencies

```bash
go mod tidy
```

---

### 4. Run the server

```bash
go run cmd/server/main.go
```

---

## 📚 API Documentation

Swagger UI is available at:

```
http://localhost:8080/swagger/index.html
```

---

## 🔐 Authentication

All protected routes require a Bearer token:

```
Authorization: Bearer YOUR_TOKEN
```

---

## 📌 Example Usage

### Get todos with pagination, filtering and sorting

```
GET /api/todos?page=1&limit=10&completed=true&sort=desc
```

---

## 📂 Project Structure

```
cmd/
  server/        -> main application entry point

internal/
  handler/       -> HTTP handlers
  service/       -> business logic
  repository/    -> database layer
  middleware/    -> auth middleware
  models/        -> database models
  utils/         -> helpers (JWT, validation, errors)
  config/        -> environment config

docs/            -> Swagger docs
```

---

## 🧠 Notes

This API uses stateless JWT authentication.

Logout is not implemented because JWT is handled client-side.
In a production system, you would typically use:

* token blacklist (e.g. Redis), or
* refresh token strategy

---

## 👨‍💻 Author

Backend project built in Go for learning and portfolio purposes.
