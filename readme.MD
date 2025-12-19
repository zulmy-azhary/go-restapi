# Go REST API

Complete REST API with JWT authentication and product CRUD using Go, Fiber, GORM, and PostgreSQL with clean architecture pattern.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [Response Format](#response-format)
- [Testing with cURL](#testing-with-curl)
- [Testing with Postman](#testing-with-postman)
- [Architecture](#architecture)
- [Database Schema](#database-schema)
- [Contributing](#contributing)
- [License](#license)
- [Support](#support)

## Features

- ✅ **JWT Authentication** - Secure token-based authentication
- ✅ **GORM ORM** - Auto migration, soft delete, associations
- ✅ **Health Check** - 3 types of health check endpoints
- ✅ **CRUD Operations** - Complete product management
- ✅ **Pagination** - Built-in pagination support
- ✅ **Search** - Case-insensitive search functionality
- ✅ **Clean Architecture** - Repository, Service, Handler pattern
- ✅ **Input Validation** - Request validation
- ✅ **Error Handling** - Consistent error responses
- ✅ **Password Hashing** - Bcrypt encryption
- ✅ **Soft Delete** - Records aren't permanently deleted
- ✅ **Authorization** - Users can only modify their own products
- ✅ **CORS Enabled** - Cross-Origin Resource Sharing support
- ✅ **Logging** - Request/Response logging middleware
- ✅ **Recovery** - Panic recovery middleware

## Tech Stack

- **Go** - Programming language
- **Fiber** - Web framework
- **GORM** - ORM library
- **PostgreSQL** - Database
- **JWT** - Authentication tokens
- **Bcrypt** - Password hashing
- **GoDotEnv** - Environment variable management

## Project Structure

```
go-rest-api/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go               # Configuration management
│   ├── database/
│   │   └── database.go             # Database connection & migration
│   ├── models/
│   │   ├── user.go                 # User model
│   │   └── product.go              # Product model
│   ├── dto/
│   │   ├── auth_dto.go             # Auth data transfer objects
│   │   └── product_dto.go          # Product data transfer objects
│   ├── repositories/
│   │   ├── user_repository.go      # User data access layer
│   │   └── product_repository.go   # Product data access layer
│   ├── services/
│   │   ├── auth_service.go         # Auth business logic
│   │   └── product_service.go      # Product business logic
│   ├── handlers/
│   │   ├── auth_handler.go         # Auth HTTP handlers
│   │   ├── product_handler.go      # Product HTTP handlers
│   │   └── health_handler.go       # Health check handlers
│   ├── middleware/
│   │   └── auth_middleware.go      # JWT authentication middleware
│   ├── routes/
│   │   └── routes.go               # Route definitions
│   └── utils/
│       ├── errors.go               # List of errors used in response.go
│       ├── jwt.go                  # JWT utilities
│       ├── password.go             # Password hashing utilities
│       └── response.go             # Response formatting utilities
├── postman/
│   └── go-rest-api.postman_collection.json
├── .env.example                    # Environment variables template
├── .gitignore                      # Git ignore rules
├── go.mod                          # Go module dependencies
└── README.md                       # Project documentation
```

## Prerequisites

Before starting, make sure you have installed:

- **Go** ([Download](https://golang.org/dl/))
- **PostgreSQL** 12 or higher ([Download](https://www.postgresql.org/download/))
- **Git** ([Download](https://git-scm.com/downloads))
- **Text Editor/IDE** (VS Code, GoLand, etc)

## Installation

### 1. Clone Repository

```bash
git clone <repository-url>
cd go-rest-api
```

### 2. Install Dependencies

```bash
go mod download
```

or

```bash
go mod tidy
```

### 3. Setup Database

Create a new PostgreSQL database:

```sql
CREATE DATABASE go_rest_api;
```

Or using psql command:

```bash
psql -U postgres
CREATE DATABASE go_rest_api;
\q
```

## Configuration

### 1. Create Environment File

Copy file `.env.example` ke `.env`:

```bash
cp .env.example .env
```

### 2. Update Environment Variables

Edit file `.env` file according to your configuration:

```env
# Server Configuration
PORT=8080
ENV=development

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password_here
DB_NAME=go_rest_api
DB_SSLMODE=disable

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-change-this
JWT_EXPIRE_MINUTES=1440
```

**⚠️ Important:**
- Replace `DB_PASSWORD` with your PostgreSQL password
- Replace `JWT_SECRET` with a secure key for production
- Do not commit the `.env` file to the repository

## Running the Application

### Development Mode

```bash
go run cmd/api/main.go
```

### Build and Run

```bash
# Build aplikasi
go build -o bin/api cmd/api/main.go

# Run aplikasi
./bin/api
```

### Build for Production

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o bin/api-linux cmd/api/main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o bin/api-windows.exe cmd/api/main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o bin/api-macos cmd/api/main.go
```

The server will run at: `http://localhost:8080`

## API Documentation

### Base URL

```
http://localhost:8080/api
```

### Endpoints Overview

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/health` | Basic health check | ❌ |
| GET | `/health/detailed` | Detailed health check | ❌ |
| GET | `/health/ready` | Readiness probe | ❌ |
| POST | `/auth/register` | Register new user | ❌ |
| POST | `/auth/login` | Login user | ❌ |
| GET | `/auth/me` | Get user profile | ✅ |
| POST | `/products` | Create product | ✅ |
| GET | `/products` | Get all products | ✅ |
| GET | `/products/:id` | Get product by ID | ✅ |
| PUT | `/products/:id` | Update product | ✅ |
| DELETE | `/products/:id` | Delete product | ✅ |

---

### Health Check Endpoints

#### 1. Basic Health Check

**GET** `/api/health`

**Response:**
```json
{
  "success": true,
  "message": "Server is running",
  "data": {
    "service": "go-rest-api",
    "status": "ok",
    "timestamp": "2025-12-17T13:03:55+08:00",
    "version": "1.0.0"
  },
}
```

#### 2. Detailed Health Check

**GET** `/api/health/detailed`

**Response:**
```json
{
  "success": true,
  "message": "Service is ready",
  "data": {
    "service": "go-rest-api",
    "status": "ok",
    "timestamp": "2025-12-17T13:05:21+08:00",
    "version": "1.0.0",
    "database": {
      "status": "up",
      "idle": 1,
      "in_use": 0,
      "max_open_conns": 0,
      "open_conns": 1
    }
  }
}
```

#### 3. Readiness Probe

**GET** `/api/health/ready`

**Response:**
```json
{
  "ready": true
}
```

---

### Authentication Endpoints

#### 1. Register

**POST** `/api/auth/register`

**Request Body:**
```json
{
  "email": "john.doe@example.com",
  "username": "johndoe",
  "password": "password123",
  "name": "John Doe"
}
```

**Validation Rules:**
- `email`: required, valid email format
- `username`: required, minimum 3 characters, unique
- `password`: required, minimum 6 characters
- `name`: required, minimum 3 characters

**Success Response (201):**
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "email": "john.doe@example.com",
      "username": "johndoe",
      "name": "John Doe",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    }
  }
}
```

**Error Response (400):**
```json
{
  "success": false,
  "message": "username already registered"
}
```

#### 2. Login

**POST** `/api/auth/login`

**Request Body:**
```json
{
  "username": "johndoe",
  "password": "password123"
}
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "email": "john.doe@example.com",
      "username": "johndoe",
      "name": "John Doe",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    }
  }
}
```

**Error Response (401):**
```json
{
  "success": false,
  "message": "invalid credentials"
}
```

#### 3. Get Profile

**GET** `/api/auth/me`

**Headers:**
```
Authorization: Bearer <YOUR_JWT_TOKEN>
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "Profile retrieved",
  "data": {
    "id": 1,
    "email": "john.doe@example.com",
    "username": "johndoe",
    "name": "John Doe",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
}
```

**Error Response (401):**
```json
{
  "success": false,
  "message": "Invalid or expired token"
}
```

**Error Response (404):**
```json
{
  "success": false,
  "message": "user not found"
}
```

---

### Product Endpoints

**⚠️ Note:** All product endpoints require JWT token in the header:
```
Authorization: Bearer <YOUR_JWT_TOKEN>
```

#### 1. Create Product

**POST** `/api/products`

**Headers:**
```
Authorization: Bearer <YOUR_JWT_TOKEN>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Laptop Gaming",
  "description": "High-end gaming laptop with RTX 4090",
  "price": 25000000,
  "stock": 15
}
```

**Validation Rules:**
- `name`: required, minimum 3 characters
- `description`: optional
- `price`: required, must be greater than 0
- `stock`: required, must be greater than or equal to 0

**Success Response (201):**
```json
{
  "success": true,
  "message": "Product created successfully",
  "data": {
    "id": 1,
    "name": "Laptop Gaming",
    "description": "High-end gaming laptop with RTX 4090",
    "price": 25000000,
    "stock": 15,
    "created_by": 1,
    "user": {
      "id": 1,
      "email": "john.doe@example.com",
      "username": "johndoe",
      "name": "John Doe",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    },
    "created_at": "2024-01-15T11:00:00Z",
    "updated_at": "2024-01-15T11:00:00Z"
  }
}
```

#### 2. Get All Products

**GET** `/api/products?page=1&perPage=10&search=laptop`

**Query Parameters:**
- `page` (optional): Page number, default = 1
- `perPage` (optional): Items per page, default = 10
- `search` (optional): Search by product name (case-insensitive)

**Success Response (200):**
```json
{
  "success": true,
  "message": "Products retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Laptop Gaming",
      "description": "High-end gaming laptop with RTX 4090",
      "price": 25000000,
      "stock": 15,
      "created_by": 1,
      "user": {
        "id": 1,
        "email": "john.doe@example.com",
        "username": "johndoe",
        "name": "John Doe",
        "created_at": "2024-01-15T10:30:00Z",
        "updated_at": "2024-01-15T10:30:00Z"
      },
      "created_at": "2024-01-15T11:00:00Z",
      "updated_at": "2024-01-15T11:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "per_page": 10,
    "total": 1,
    "total_pages": 1
  }
}
```

#### 3. Get Product by ID

**GET** `/api/products/:id`

**Success Response (200):**
```json
{
  "success": true,
  "message": "Product retrieved successfully",
  "data": {
    "id": 1,
    "name": "Laptop Gaming",
    "description": "High-end gaming laptop with RTX 4090",
    "price": 25000000,
    "stock": 15,
    "created_by": 1,
    "user": {
      "id": 1,
      "email": "john.doe@example.com",
      "username": "johndoe",
      "name": "John Doe",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    },
    "created_at": "2024-01-15T11:00:00Z",
    "updated_at": "2024-01-15T11:00:00Z"
  }
}
```

**Error Response (404):**
```json
{
  "success": false,
  "message": "product not found"
}
```

#### 4. Update Product

**PUT** `/api/products/:id`

**Headers:**
```
Authorization: Bearer <YOUR_JWT_TOKEN>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Laptop Gaming Pro",
  "description": "Updated description",
  "price": 27000000,
  "stock": 20
}
```

**Note:** All fields are optional. Only the sent fields will be updated.

**Success Response (200):**
```json
{
  "success": true,
  "message": "Product updated successfully",
  "data": {
    "id": 1,
    "name": "Laptop Gaming Pro",
    "description": "Updated description",
    "price": 27000000,
    "stock": 20,
    "created_by": 1,
    "user": {
      "id": 1,
      "email": "john.doe@example.com",
      "username": "johndoe",
      "name": "John Doe",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    },
    "created_at": "2024-01-15T11:00:00Z",
    "updated_at": "2024-01-15T12:00:00Z"
  }
}
```

**Error Response (403):**
```json
{
  "success": false,
  "message": "unauthorized to update this product"
}
```

#### 5. Delete Product

**DELETE** `/api/products/:id`

**Headers:**
```
Authorization: Bearer <YOUR_JWT_TOKEN>
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "Product deleted successfully"
}
```

**Error Response (403):**
```json
{
  "success": false,
  "message": "unauthorized to delete this product"
}
```

**Note:** This is a soft delete, data is not actually deleted from the database.

---

## Response Format

### Success Response
All successful responses in this application return a structure like this:
```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... }
}
```
Or it can also be without returning `data`, like this :
```json
{
  "success": true,
  "message": "Operation successful"
}
```

### Error Response
All error responses in this application return a JSON structure like this:
```json
{
  "success": false,
  "message": "Error message here",
}
```
Can also include error details like this:
```json
{
  "success": false,
  "message": "Error message here",
  "error": {
    "code": 400,
    "type": "BAD_REQUEST"
  }
}
```

### Paginated Response

```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [ ... ],
  "pagination": {
    "page": 1,
    "per_page": 10,
    "total": 100,
    "total_pages": 10
  }
}
```

---

## Testing with cURL

### 1. Register User

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "username": "johndoe",
    "password": "password123",
    "name": "John Doe"
  }'
```

### 2. Login

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "password": "password123"
  }'
```

**Save the token from response!**

### 3. Create Product

```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <YOUR_TOKEN_HERE>" \
  -d '{
    "name": "Laptop Gaming",
    "description": "High-end gaming laptop",
    "price": 25000000,
    "stock": 15
  }'
```

### 4. Get All Products

```bash
curl -X GET "http://localhost:8080/api/products?page=1&perPage=10" \
  -H "Authorization: Bearer <YOUR_TOKEN_HERE>"
```

### 5. Get Product by ID

```bash
curl -X GET http://localhost:8080/api/products/1 \
  -H "Authorization: Bearer <YOUR_TOKEN_HERE>"
```

### 6. Update Product

```bash
curl -X PUT http://localhost:8080/api/products/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <YOUR_TOKEN_HERE>" \
  -d '{
    "name": "Updated Laptop Gaming",
    "price": 27000000,
    "stock": 20
  }'
```

### 7. Delete Product

```bash
curl -X DELETE http://localhost:8080/api/products/1 \
  -H "Authorization: Bearer <YOUR_TOKEN_HERE>"
```

### 8. Search Products

```bash
curl -X GET "http://localhost:8080/api/products?search=laptop" \
  -H "Authorization: Bearer <YOUR_TOKEN_HERE>"
```

---

## Testing with Postman
### Prerequisites
- Install [Postman Desktop](https://www.postman.com/downloads/) or use [Postman Web](https://web.postman.com/)
- Ensure the API application is running at `http://localhost:8080`


### Import Postman Collection

#### Method 1: Import dari File

1. Download the JSON file available [here](https://github.com/zulmy-azhary/go-restapi/blob/master/postman/go-rest-api.postman_collection.json)
2. Open Postman
3. Select the **File** tab
4. Click **Import...** in the top left corner
5. **Drag & drop** the file or click **Choose Files**
6. Click **Import**

#### Method 2: Import dari Url

1. Open Postman
2. Select the **File** tab
3. Click **Import...** in the top left corner
4. Enter this [URL](https://raw.githubusercontent.com/zulmy-azhary/go-restapi/refs/heads/master/postman/go-rest-api.postman_collection.json) in the input field
5. Click **Import**

After importing, you will see a folder structure like this:
```
📁 Go REST API - Complete Collection
├── 📁 Health Check
│   ├── GET Basic Health Check
│   ├── GET Detailed Health Check
│   └── GET Readiness Probe
├── 📁 Authentication
│   ├── POST Register User
│   ├── POST Login
│   └── GET Get Profile
├── 📁 Products
│   ├── POST Create Product
│   ├── GET Get All Products
│   ├── GET Get Products with Search
│   ├── GET Get Product by ID
│   ├── PUT Update Product
│   └── DELETE Delete Product
└── 📁 Test Scenarios
    ├── POST Register Another User
    ├── PUT Try Update Other User's Product (Should Fail)
    ├── GET Try Access Without Token (Should Fail)
    ├── POST Try Register with Duplicate Email (Should Fail)
    ├── POST Try Register with Duplicate Username (Should Fail)
    └── POST Try Login with Wrong Password (Should Fail)
```
---

## Architecture

### Clean Architecture Pattern

This application uses clean architecture with separation of concerns:

```
┌────────────────────────────────────────┐
│           HTTP Layer (Fiber)           │
│   ┌────────────────────────────────┐   │
│   │            Handlers            │   │
│   │    (HTTP Request/Response)     │   │
│   └───────────────┬────────────────┘   │
└───────────────────┼────────────────────┘
                    │
┌───────────────────▼────────────────────┐
│          Business Logic Layer          │
│   ┌────────────────────────────────┐   │
│   │            Services            │   │
│   │    (Business Rules & Logic)    │   │
│   └───────────────┬────────────────┘   │
└───────────────────┼────────────────────┘
                    │
┌───────────────────▼────────────────────┐
│            Data Access Layer           │
│   ┌────────────────────────────────┐   │
│   │          Repositories          │   │
│   │      (Database Operations)     │   │
│   └───────────────┬────────────────┘   │
└───────────────────┼────────────────────┘
                    │
┌───────────────────▼────────────────────┐
│             Database Layer             │
│   ┌────────────────────────────────┐   │
│   │        PostgreSQL + GORM       │   │
│   └────────────────────────────────┘   │
└────────────────────────────────────────┘
```

### Layer Responsibilities

1. **Handlers**: Receive HTTP requests, validate input, call services, return responses
2. **Services**: Business logic, validate business rules, orchestration
3. **Repositories**: Data access, query database, CRUD operations
4. **Models**: Database entities dan schema

### Advantages

- ✅ **Testability**: Each layer can be tested independently
- ✅ **Maintainability**: Easy to maintain and modify
- ✅ **Scalability**: Easy to scale and extend
- ✅ **Flexibility**: Easy to change implementation (e.g., change database)

---

## Database Schema

### Users Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_users_deleted_at ON users(deleted_at);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);
```

### Products Table

```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(15, 2) NOT NULL,
    stock INTEGER NOT NULL DEFAULT 0,
    created_by INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_products_deleted_at ON products(deleted_at);
CREATE INDEX idx_products_name ON products(name);
CREATE INDEX idx_products_created_by ON products(created_by);
```

### Entity Relationship

```
┌─────────────────┐          ┌─────────────────┐
│     Users       │          │    Products     │
├─────────────────┤          ├─────────────────┤
│ id (PK)         │◄────────┐│ id (PK)         │
│ email           │         ││ name            │
│ username        │         ││ description     │
│ password        │         ││ price           │
│ name            │         ││ stock           │
│ created_at      │         └│ created_by (FK) │
│ updated_at      │          │ created_at      │
│ deleted_at      │          │ updated_at      │
└─────────────────┘          │ deleted_at      │
                             └─────────────────┘
```

---

## Security Features

### 1. Password Security
- Passwords are hashed using **bcrypt** with cost factor 10
- Passwords are never displayed in JSON responses

### 2. JWT Authentication
- Token expires in 24 hours (configurable)
- Token uses HS256 signing method
- Secret key must be changed for production

### 3. Authorization
- Users can only update/delete products they created
- Middleware verifies token on every protected route

### 4. SQL Injection Prevention
- GORM uses prepared statements
- Automatic parameter binding

### 5. CORS
- CORS enabled with customizable configuration
- Default: allow all origins (for development)

---

## Error Handling

### HTTP Status Codes

- `200 OK`: Request successful
- `201 Created`: Resource successfully created
- `400 Bad Request`: Invalid request body or validation error
- `401 Unauthorized`: Authentication failed or invalid token
- `403 Forbidden`: User does not have access
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

### Common Errors

#### 1. Missing Authorization Header
```json
{
  "success": false,
  "message": "Missing authorization header"
}
```

#### 2. Invalid Token
```json
{
  "success": false,
  "message": "Invalid or expired token"
}
```

#### 3. Validation Error
```json
{
  "success": false,
  "message": "Invalid request body"
}
```

---

## Troubleshooting

### Database Connection Error

**Error:** `failed to connect to database`

**Solution:**
1. Ensure PostgreSQL is running
2. Check credentials in `.env` file
3. Ensure database has been created
4. Test connection: `psql -U postgres -d go_rest_api`

### Port Already in Use

**Error:** `bind: address already in use`

**Solution:**
1. Change port in `.env`: `PORT=8081`
2. Or kill the process using the port:
   ```bash
   # Linux/Mac
   lsof -ti:8080 | xargs kill -9
   
   # Windows
   netstat -ano | findstr :8080
   taskkill /PID <PID> /F
   ```

### Migration Error

**Error:** `failed to migrate database`

**Solution:**
1. Drop existing tables:
   ```sql
   DROP TABLE products;
   DROP TABLE users;
   ```
2. Restart application

### Token Expired

**Error:** `Invalid or expired token`

**Solution:**
1. Login again to get a new token
2. Token is valid for 24 hours

---

## Performance Tips

### 1. Database Indexing
- Indexes already added on foreign keys
- Index on `deleted_at` for soft delete queries
- Index on `email` for faster login

### 2. Pagination
- Always use pagination for list queries
- Default per page: 10, max recommended: 100

### 3. Connection Pooling
- GORM uses connection pooling by default
- Can be configured in `database.go`

### 4. Caching (Optional)
- Add Redis for caching frequent queries
- Cache JWT token validation results

---

## Deployment

### Deploy to Production

1. **Build the application:**
   ```bash
   CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/api/main.go
   ```

2. **Setup environment variables**

3. **Run migrations**

4. **Start the application:**
   ```bash
   ./main
   ```

### Docker Deployment (Optional)

**Dockerfile:**
```dockerfile
FROM golang:1.25.1-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
EXPOSE 8080
CMD ["./main"]
```

**Build & Run:**
```bash
docker build -t go-rest-api .
docker run -p 8080:8080 go-rest-api
```

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## License

This project is licensed under the MIT License.

---

## Support

If you have any questions or issues:

1. Open an issue on the GitHub repository
2. Email: zulmyazhary32@gmail.com

---

**Built with ❤️ using Go, Fiber, GORM, and PostgreSQL**