# CodeQL Test Repo - Go REST API

A simple Go REST API built with **Gin** and **SQLite** for testing CodeQL security analysis.

## Project Structure

```
.
├── main.go              # Application entrypoint and route setup
├── database/
│   └── db.go            # SQLite initialization and schema
├── handlers/
│   ├── auth.go          # Register and login endpoints
│   ├── user.go          # Profile fetch endpoint
│   ├── search.go        # User search endpoint
│   └── file.go          # File download endpoint
├── middleware/
│   └── auth.go          # JWT authentication middleware
├── models/
│   └── user.go          # Data models and request types
├── uploads/
│   └── sample.txt       # Sample downloadable file
├── Dockerfile
└── README.md
```

## Setup

### Prerequisites

- Go 1.21+
- GCC (for SQLite CGO compilation)

### Run Locally

```bash
go mod tidy
go run .
```

The server starts on `http://localhost:8080`.

### Run with Docker

```bash
docker build -t codeql-test-api .
docker run -p 8080:8080 codeql-test-api
```

## API Endpoints

| Method | Path           | Auth     | Description          |
|--------|----------------|----------|----------------------|
| POST   | /register      | No       | Register a new user  |
| POST   | /login         | No       | Login, get JWT token |
| GET    | /profile/:id   | Bearer   | Get user profile     |
| GET    | /search?q=     | Bearer   | Search users         |
| GET    | /download?file=| Bearer   | Download a file      |

## Sample curl Commands

### Register a user
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","email":"alice@example.com","password":"securepass123"}'
```

### Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"securepass123"}'
```

### Get profile (use token from login response)
```bash
curl http://localhost:8080/profile/1 \
  -H "Authorization: Bearer <your-token>"
```

### Search users
```bash
curl "http://localhost:8080/search?q=alice" \
  -H "Authorization: Bearer <your-token>"
```

### Download a file
```bash
curl "http://localhost:8080/download?file=sample.txt" \
  -H "Authorization: Bearer <your-token>" \
  --output sample.txt
```

## Branches

- **main** - Clean, secure implementation
- **feature/** - Branches with intentional vulnerabilities for CodeQL testing
