# 🔐 SentinelAuth

**SentinelAuth** is a modern authentication and authorization microservice built with **Golang**, designed to be secure, scalable, and modular. It features OTP, email verification, JWT-based access/refresh tokens, role-based access control (RBAC), and complete Swagger/OpenAPI documentation.

---

## 🚀 Features

- ✅ User Registration & Login
- ✅ JWT Access + Refresh Tokens
- ✅ Logout & Token Revocation (Blacklist with Redis)
- ✅ OTP Verification (via simulated email/SMS)
- ✅ Email Verification with expiring tokens
- ✅ Role Management (CRUD)
- ✅ Role Assignment to Users
- ✅ Middleware-based RBAC
- ✅ Pagination for GET endpoints
- ✅ Swagger Documentation (`/swagger/index.html`)

---

## 📦 Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gin
- **Database:** PostgreSQL (via GORM)
- **Cache:** Redis
- **Docs:** Swagger via Swag CLI
- **Architecture:** Clean Architecture & Modular

---

## 📁 Project Structure

SentinelAuth/
├── cmd/ # App entrypoint (main.go)
├── internal/
│ ├── config/ # Environment config
│ ├── domain/ # Entities & repository interfaces
│ ├── usecase/ # Business logic
│ ├── handler/ # HTTP handlers (controllers)
│ ├── repository/ # GORM & Redis implementations
│ ├── middleware/ # JWT & RBAC middleware
│ └── utils/ # Helper functions (token, hash, etc.)
├── docs/ # Swagger generated files
├── Dockerfile
├── docker-compose.yml
├── .env
├── go.mod
└── go.sum


---

## ⚙️ Environment Variables (`.env`)

```env
PORT=8080
DB_DSN=host=localhost user=postgres password=yourpass dbname=sentinel port=5432 sslmode=disable
REDIS_ADDR=localhost:6379
JWT_SECRET=your_super_secret_key

🧪 Run Locally

# 1. Generate Swagger docs
swag init --generalInfo cmd/main.go --output docs --parseDependency --parseInternal

# 2. Start the server
go run cmd/main.go

# 3. Access Swagger docs
http://localhost:8080/swagger/index.html

🐳 Docker Support

docker-compose up --build

    You can define Dockerfile and docker-compose.yml accordingly to run PostgreSQL, Redis, and the app together.

🔒 Role-Based Access Control (RBAC)

    Users can be assigned to roles.

    Middleware like RoleRequired("admin") ensures access control per endpoint.

✅ Swagger Annotation Example

// @Summary Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse
// @Router /login [post]

📬 Contact & Contribution

Feel free to fork the repo or suggest improvements via issues or PRs.
