# Gin API Template

A modern backend API development template based on the Go-Gin framework, integrated with GORM, automatic parameter binding, testing framework, and more.

## 🚀 Features

- **Gin Framework**: High-performance HTTP web framework
- **GORM Integration**: Powerful ORM library supporting multiple databases
- **Automatic Parameter Binding**: Automatic parsing and validation of GET/POST parameters
- **Testing Framework**: Integrated testify for unit and integration testing
- **Middleware Support**: CORS, parameter validation, and other middleware
- **Unified Response Format**: Standardized API response structure
- **Layered Architecture**: Clear separation of Controller -> Service -> Model layers
- **Structured Logging**: High-performance logging with zerolog
- **Swagger Documentation**: Auto-generated interactive API documentation

## 📁 Project Structure

```
gin-template/
├── cmd/
│   └── server/
│       └── main.go        # Application entry point
├── go.mod                 # Go module file
├── config/                # Configuration management
├── database/              # Database connection
├── models/                # Data models (User only)
├── middleware/            # Middleware (smart parameter binding)
├── service/               # Business logic layer
├── controller/            # Controller layer (new architecture)
├── router/                # Route configuration
├── docs/                  # Swagger documentation
├── test/                  # Test files
└── examples/              # API examples
```

## 🛠️ Quick Start

### 1. Clone the Repository

```bash
git clone <repository-url>
cd gin-template
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Configure Environment Variables (Optional)

```bash
export PORT=8080
export DB_DRIVER=sqlite
export DB_DSN=test.db
export LOG_LEVEL=info
export LOG_FORMAT=pretty
```

### 4. Run the Project

```bash
# Using Makefile
make run

# Or run directly
go run cmd/server/main.go
```

The server will start at `http://localhost:8080`

### 5. Run Tests

```bash
go test ./test/...
```

## 📚 API Documentation

### Swagger UI

After starting the server, visit the following addresses to view API documentation:

- **Swagger UI**: <http://localhost:8080/swagger/index.html>
- **JSON Documentation**: <http://localhost:8080/swagger/doc.json>
- **YAML Documentation**: <http://localhost:8080/swagger/doc.yaml>

### User Endpoints

#### Create User

```http
POST /api/v1/users
Content-Type: application/json

{
    "name": "John Doe",
    "email": "john@example.com",
    "age": 25,
    "phone": "1234567890"
}
```

#### Get User List

```http
GET /api/v1/users?page=1&page_size=10&name=John&email=john
```

#### Get Single User

```http
GET /api/v1/users/{id}
```

#### Update User

```http
PUT /api/v1/users/{id}
Content-Type: application/json

{
    "name": "Jane Doe",
    "age": 30
}
```

#### Delete User

```http
DELETE /api/v1/users/{id}
```

### Health Check

```http
GET /health
```

## 🏗️ Architecture Overview

### Automatic Parameter Binding

The project uses custom middleware for automatic parameter binding:

```go
// Bind JSON request body
middleware.BindJSON(&models.CreateUserRequest{})

// Bind query parameters
middleware.BindQuery(&models.GetUsersQuery{})
```

### Unified Response Format

All API responses use a unified format:

```json
{
    "code": 200,
    "message": "success",
    "data": { ... }
}
```

### Data Validation

Uses validator tags for data validation:

```go
type CreateUserRequest struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"min=1,max=150"`
}
```

## 🚀 Enhanced Features

### Structured Logging (Zerolog)

- ✅ **High Performance**: Based on [zerolog](https://github.com/rs/zerolog) structured logging
- ✅ **Flexible Configuration**: Support for different log levels and output formats
- ✅ **Developer Friendly**: Colorful console output for development environment
- ✅ **Production Ready**: JSON format output for production environment

### Swagger API Documentation

- ✅ **Auto Generation**: Automatically generate API documentation from code comments
- ✅ **Interactive Interface**: Support online API testing
- ✅ **Standard Format**: Compliant with OpenAPI 3.0 specification
- ✅ **Real-time Updates**: Automatically update documentation when code changes

## 🔧 Environment Variables

```bash
# Logging configuration
export LOG_LEVEL=debug          # trace, debug, info, warn, error, fatal, panic
export LOG_FORMAT=pretty        # pretty/console (development) or json (production)

# Server configuration
export PORT=8080
export GIN_MODE=release         # Set for production environment

# Database configuration
export DB_DRIVER=sqlite
export DB_DSN=test.db
```

## 🛠️ Development Commands

```bash
# Generate Swagger documentation
make docs

# Format Swagger annotations
make docs-fmt

# Run project (auto-generate documentation)
make run

# Run tests
make test

# Code formatting
make fmt

# Code linting
make lint

# Build project
make build

# Development mode with auto-reload
make dev
```

## 🧪 Testing

The project includes a comprehensive test suite:

- **Unit Tests**: Test individual component functionality
- **Integration Tests**: Test complete API flows
- **Test Utilities**: Provide convenient test helper functions

Run specific tests:

```bash
# Run user tests
go test ./test -run TestUser

# Show detailed output
go test ./test -v

# Run tests with coverage
make test-cover
```

## 🔧 Extension Features

### Adding New Models

1. Create new model files in `models/`
2. Add AutoMigrate in `database/database.go`
3. Create corresponding Service and Controller
4. Add routes in `router/router.go`

### Adding Middleware

Create new middleware in `middleware/`, then use it in routes:

```go
r.Use(middleware.YourCustomMiddleware())
```

### Database Configuration

Supports multiple databases:

- SQLite (default)
- MySQL
- PostgreSQL

Configure via environment variables `DB_DRIVER` and `DB_DSN`.

## 🐳 Docker Support

```bash
# Build Docker image
make docker-build

# Run with Docker
make docker-run
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📞 Support

If you have any questions or need help, please create an issue on GitHub.

---

**Happy Coding! 🚀**
