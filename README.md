# ecom

[![Go Version](https://img.shields.io/badge/Go-1.25-blue.svg)](https://go.dev/)
[![Status](https://img.shields.io/badge/status-work--in--progress-orange.svg)]()

> **This is a personal backend development practice project.** I'm building an e-commerce API from scratch to sharpen my skills in Go, REST API design, database modeling, authentication, and production-ready backend patterns. It is **a work in progress** — many features are still being implemented and the codebase will evolve as I learn.

---

## About

**ecom** is a backend API for an e-commerce platform. The goal is to implement core features such as user authentication, product catalog, shopping cart, orders, and payments — all while following clean architecture principles and good engineering practices.

This repository is where I'm actively learning and experimenting. Expect frequent refactors, new features, and improvements over time.

---

## Tech Stack

- **Language:** [Go 1.25](https://go.dev/)
- **HTTP Router:** [gorilla/mux](https://github.com/gorilla/mux)
- **Database:** MySQL
- **Database Driver:** [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- **Configuration:** Environment variables via [godotenv](https://github.com/lpernett/godotenv)
- **Password Hashing:** [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) (from `golang.org/x/crypto`)
- **Payload Validation:** [go-playground/validator](https://github.com/go-playground/validator)
- **Containerization:** Docker & Docker Compose (for MySQL)
- **Build Tool:** Makefile

---

## Project Structure

```
.
├── cmd/
│   ├── api/              # HTTP server setup and routing
│   └── main.go           # Application entry point
├── config/               # Environment configuration
├── db/                   # Database connection setup
├── service/              # Business logic organized by domain
│   ├── auth/             # Password hashing utilities
│   └── user/             # User service (login, register, store, tests)
├── types/                # Shared domain types and interfaces
├── utils/                # HTTP request/response helpers
├── bin/                  # Compiled binary output
├── compose.yaml          # Docker Compose for local MySQL
├── Makefile              # Common build and run commands
└── README.md             # You are here
```

---

## Current Status

The project is in its very early stages. Here's what exists so far:

| Feature | Status |
|---------|--------|
| Project setup and HTTP server | Done |
| Environment configuration | Done |
| MySQL connection via Docker Compose | Done |
| Shared types, utils, and repository interfaces | Done |
| Payload validation with `go-playground/validator` | Done |
| User registration handler with bcrypt password hashing | Done |
| User store/repository scaffolding | Done |
| Unit tests for user handlers (valid and invalid payloads) | Done |
| User login endpoint | In Progress |
| JWT-based authentication | Planned |
| Product catalog | Planned |
| Shopping cart | Planned |
| Orders & payments | Planned |
| Integration tests & API documentation | Planned |

---

## Getting Started

### Prerequisites

- [Go 1.25+](https://go.dev/dl/)
- [Docker](https://www.docker.com/) & Docker Compose
- Make (optional, for convenience commands)

### 1. Clone the repository

```bash
git clone https://github.com/vhs0s0/ecom.git
cd ecom
```

### 2. Set up environment variables

Create a `.env` file in the project root based on `.env.example`:

```bash
cp .env.example .env
```

Example `.env`:

```env
PUBLIC_HOST=http://localhost
PORT=8080

DB_USER=root
DB_PASSWORD=root
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=ecom
```

### 3. Start the database

```bash
docker compose up -d
```

This will start a MySQL container on port `3306`.

### 4. Run the application

Using Make:

```bash
make run
```

Or directly with Go:

```bash
go run cmd/main.go
```

The server will start at `http://localhost:8080`.

---

## API Endpoints

Base URL: `http://localhost:8080/api/v1`

| Method | Endpoint | Description | Status |
|--------|----------|-------------|--------|
| `POST` | `/api/v1/login` | User login | In Progress |
| `POST` | `/api/v1/register` | User registration (validates payload, checks existing email, hashes password) | Done* |

> More endpoints will be added as the project grows.
>
> *`Done*` means the endpoint logic, validation, and password hashing are implemented. Database persistence via `CreateUser` is still being finished.

---

## Testing

```bash
make test
```

---

## Roadmap

- [x] Bootstrap Go project
- [x] Set up HTTP server and routing
- [x] Connect to MySQL with Docker Compose
- [x] Add shared types, utils, and repository interfaces
- [x] Add payload validation with `go-playground/validator`
- [x] Implement user registration handler with password hashing
- [x] Add unit tests for user handlers (valid and invalid payloads)
- [ ] Implement user login and JWT authentication
- [ ] Persist new users in the database (`CreateUser`)
- [ ] Build product catalog service
- [ ] Implement shopping cart
- [ ] Create order management
- [ ] Add integration tests and API documentation
- [ ] Deploy to a cloud provider

---

## Why This Project?

I'm using this project as a hands-on way to learn backend engineering concepts in Go. It's a safe space to experiment, make mistakes, and improve. If you're learning too, feel free to follow along or suggest improvements.

---

## Connect

- GitHub: [@vhs0s0](https://github.com/vhs0s0)
- Built as a learning project.
