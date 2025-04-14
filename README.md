# User Service

A microservice for user management built with Go and MongoDB.

## Features

- User CRUD operations (Create, Read, Update, Delete)
- MongoDB integration for data persistence
- RESTful API with JSON responses
- Input validation for email, password, and username
- Environment-based configuration
- Health check endpoint

## Prerequisites

- Go 1.18 or later
- MongoDB 4.4 or later

## Configuration

The service can be configured using environment variables:

```env
PORT=8081                            # Server port (default: 8081)
MONGO_URI=mongodb://localhost:27017  # MongoDB connection URI
DB_NAME=user_service                 # MongoDB database name
```

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd user-service
```

2. Install dependencies:
```bash
go mod download
```

3. Run the service:
```bash
go run main.go
```

## API Endpoints

### Health Check
```
GET /health
Response: 200 OK
```

### Create User
```
POST /api/v1/users
Content-Type: application/json

Request:
{
    "email": "user@example.com",
    "password": "password123",
    "username": "johndoe"
}

Response: 201 Created
{
    "id": "...",
    "email": "user@example.com",
    "username": "johndoe",
    "created_at": "...",
    "updated_at": "..."
}
```

### Get User
```
GET /api/v1/users/{id}
Response: 200 OK
```

### Get All Users
```
GET /api/v1/users
Response: 200 OK
```

### Update User
```
PUT /api/v1/users/{id}
Content-Type: application/json

Request:
{
    "email": "newemail@example.com",
    "password": "newpassword123",
    "username": "newusername"
}

Response: 200 OK
```

### Delete User
```
DELETE /api/v1/users/{id}
Response: 204 No Content
```

## Error Responses

The API returns appropriate HTTP status codes and error messages:

- 400 Bad Request: Invalid input data
- 404 Not Found: User not found
- 409 Conflict: User already exists
- 500 Internal Server Error: Server-side error

## Project Structure

```
.
├── config/         # Configuration management
├── handlers/       # HTTP request handlers
├── models/         # Data models and DTOs
├── repository/     # Data access layer
│   └── mongodb/   # MongoDB implementation
├── service/        # Business logic
└── main.go        # Application entry point
```

## License

MIT
