# task manager Service

A microservice for task management built with Go and MongoDB.

## Features

- User CRUD operations (Create, Read, Update, Delete)
- MongoDB integration for data persistence
- RESTful API with JSON responses
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
DB_NAME=task_manager                # MongoDB database name
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

## Error Responses

The API returns appropriate HTTP status codes and error messages:

- 400 Bad Request: Invalid input data
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
