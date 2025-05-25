# Golang User Management API

A RESTful API develop with Golang, Gin framework, and MongoDB for user management with JWT authentication.

## Features

- User Register and Authentication
- JWT authentication
- CRUD for users
- Password hashing with bcrypt
- MongoDB database 
- Middleware for logging and authentication
- Docker containerization
- Unit testing with mock database
- Background worker for user count monitoring

## Tech Stack

- **Language**: Golang
- **Framework**: Gin (HTTP web framework)
- **Database**: MongoDB
- **Authentication**: JWT (JSON Web Tokens)
- **Containerization**: Docker & Docker Compose

## Project Structure

```
├── config/
│   └── db.go              # Database configuration
├── controllers/
│   ├── auth.go            # Authentication controllers (register, login)
│   ├── user.go            # User CRUD controllers
│   └── user_test.go       # User Unit tests
├── middleware/
│   ├── auth.go            # JWT authentication middleware
│   └── logger.go          # Request logging middleware
├── models/
│   └── user.go            # User models and structs
├── routes/
│   └── routes.go          # API Endpoint route 
├── utils/
│   └── jwt.go             # JWT utility functions
├── workers/
│   └── user_counter.go    # Background worker
├── docker-compose.yml     # Docker services configuration
├── Dockerfile             # Container build instructions
├── go.mod                 # Go module dependencies
└── main.go                # Main Application 
```

## Setup and Installation

### Prerequisites

- Golang 1.24 or higher
- MongoDB 
- Docker and Docker Compose 

### Method 1: Local Development Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd golang-challenge
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run the application**
   ```bash
   go run main.go
   ```

### Method 2: Docker Setup

  1. **Run with Docker Compose**
   ```bash
   docker-compose up --build
   ```

This will start both MongoDB and the API server automatically.


## Database Configuration

### Local MongoDB
```
mongodb://localhost:27017/golang_challenge
```

### Docker MongoDB (from docker-compose.yml)
```
mongodb://usernamedb:dbpassword@mongodb:27017/golang_challenge?authSource=admin
```

## API Endpoints

### Public Endpoints (No Authentication Required)

#### Health Check
```http
GET /api/healthcheck
```

**Response:**
```json
{
  "status": "healthy",
  "service": "golang-api",
  "time": "2025-XX-XX"
}
```

#### User Registration
```http
POST /api/register
Content-Type: application/json

{
  "name": "mosu",
  "email": "mosu@gmail.com",
  "password": "12345678",
  "confirm_password": "12345678"
}
```

**Response:**
```json
{
  "message": "User created successfully",
  "user": {
    "id": "60f7b3b4b3b4b3b4b3b4b3b4",
    "name": "mosu",
    "email": "mosu@gmail.com",
    "created_at": "2025-XX-XX"
  }
}
```

#### User Login
```http
POST /api/login
Content-Type: application/json

{
  "email": "mosu@gmail.com",
  "password": "12345678"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "60f7b3b4b3b4b3b4b3b4b3b4",
    "name": "mosu",
    "email": "mosu@gmail.com",
    "created_at": "2025-XX-XX"
  }
}
```

### Protected Endpoints (JWT Authentication Required)

All protected endpoints require the `Authorization` header with Bearer token:
```
Authorization: Bearer <your-jwt-token>
```

#### Create User
```http
POST /api/users
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "mosuAdd",
  "email": "mosuadd@gmail.com",
  "password": "12345678",
  "confirm_password": "12345678"
}
```

**Response:**
```json
[
{
    "id": "6833634974c11a0a92a66fbf",
    "name": "mosuAdd",
    "email": "mosuadd@gmail.com",
    "created_at": "2025-05-25T18:36:57.270356844Z"
}
]
```

#### Get User by ID
```http
GET /api/users/{user_id}
Authorization: Bearer <token>
```

**Response:**
```json
[
  {
    "id": "68334cf674c11a0a92a66fbd",
    "name": "adminname",
    "email": "admin@gmail.com",
    "created_at": "2025-XX-XX"
}
]
```

#### Get All Users
```http
GET /api/users
Authorization: Bearer <token>
```

**Response:**
```json
[
  {
    "id": "60f7b3b4b3b4b3b4b3b4b3b4",
    "name": "mosuAdd",
    "email": "mosuadd@gmail.com",
    "created_at": "2025-XX-XX"
  }
]

[
  {
    "id": "68334cf674c11a0a92a66fbd",
    "name": "adminname",
    "email": "admin@gmail.com",
    "created_at": "2025-XX-XX"
}
]
```

#### Update User
```http
PUT /api/users/{user_id}
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "adminUpdate",
  "email": "adminname.updated@gmail.com"
}
```

**Response:**
```json
[
  {
    "id": "68334cf674c11a0a92a66fbd",
    "name": "adminUpdate",
    "email": "adminname.updated@gmail.com",
    "created_at": "2025-XX-XX"
}
]
```

#### Delete User
```http
DELETE /api/users/{user_id}
Authorization: Bearer <token>
```

**Response:**
```json
{
  "message": "User deleted successfully"
}
```

## JWT Token Usage Guide

### Token Generation
- Tokens are generated upon successful login
- Default expiration: 24 hours
- Contains user ID and email in payload

### Using JWT Tokens
1. **Login** to receive your JWT token
2. **Include** the token in the Authorization header for protected routes:
   ```
   Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
   ```

### Token Structure
```json
{
  "user_id": "60f7b3b4b3b4b3b4b3b4b3b4",
  "email": "user@example.com",
  "iat": 1673037056,
  "exp": 1673123456 
}
```

## Sample API Requests with cURL

### Register a new user
```bash
curl -X POST http://127.0.0.1:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123",
    "confirm_password": "password123"
  }'
```

### Login
```bash
curl -X POST http://127.0.0.1:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

### Get all users (with JWT token)
```bash
curl -X GET http://127.0.0.1:8080/api/users \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

## Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run specific test file
go test ./controllers -v
```

## Assumptions and Design Decisions

1. **Email as Username**: Email addresses are used as unique identifiers for login
2. **Password Requirements**: Minimum 6 characters (enforced by validation)
3. **JWT Expiration**: 24-hour token lifetime for security balance
4. **Database Choice**: MongoDB for flexibility with user data structures
5. **Error Handling**: Consistent JSON error responses across all endpoints
6. **Testing Strategy**: Unit tests with mock database for business logic testing
7. **Logging**: Request logging for monitoring and debugging
8. **Docker Support**: Full containerization for easy deployment


## API Response Codes

| Status Code | Description |
|-------------|-------------|
| 200 | OK - Request successful |
| 201 | Created - Resource created successfully |
| 400 | Bad Request - Invalid request data |
| 401 | Unauthorized - Authentication required |
| 404 | Not Found - Resource not found |
| 409 | Conflict - Resource already exists |
| 500 | Internal Server Error - Server error |
