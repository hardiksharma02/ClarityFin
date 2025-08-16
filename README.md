# ClarityFin Backend API

A robust Go backend for the ClarityFin financial management application, built with clean architecture principles, Gin, GORM, and PostgreSQL.

## 🏗️ Architecture

This project follows **Clean Architecture** principles with clear separation of concerns:

- **Domain Layer**: Business entities and interfaces
- **Repository Layer**: Data access abstraction
- **Service Layer**: Business logic implementation
- **Use Case Layer**: Application orchestration
- **Handler Layer**: HTTP request/response handling
- **Infrastructure Layer**: External concerns (database, config)

## 🚀 Features

- **Clean Architecture**: Well-structured, maintainable, and testable codebase
- **User Authentication**: JWT-based authentication with phone number and password
- **Database Integration**: PostgreSQL with GORM ORM
- **Configuration Management**: YAML-based configuration with Viper
- **RESTful API**: Clean API endpoints with proper HTTP status codes
- **Security**: Password hashing with bcrypt and JWT token validation
- **Dependency Injection**: Proper dependency management and inversion of control
- **Standardized Responses**: Consistent API response format
- **CORS Support**: Cross-origin resource sharing enabled

## 📁 Project Structure

```
clarityfin-api/
├── cmd/
│   └── api/
│       └── main.go                    # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go                  # Configuration management
│   ├── database/
│   │   └── database.go                # Database connection and setup
│   ├── domain/
│   │   ├── user.go                    # User domain entity and interfaces
│   │   └── subscription.go            # Subscription domain entity and interfaces
│   ├── dto/
│   │   ├── auth.go                    # Authentication DTOs
│   │   └── subscription.go            # Subscription DTOs
│   ├── handlers/
│   │   ├── auth_handler.go            # Authentication HTTP handlers
│   │   └── subscription_handler.go    # Subscription HTTP handlers
│   ├── middleware/
│   │   └── auth.go                    # JWT authentication middleware
│   ├── repository/
│   │   ├── user_repository.go         # User data access layer
│   │   └── subscription_repository.go # Subscription data access layer
│   └── service/
│       ├── user_service.go            # User business logic
│       ├── user_usecase.go            # User application logic
│       ├── subscription_service.go    # Subscription business logic
│       └── subscription_usecase.go    # Subscription application logic
├── pkg/
│   └── response/
│       └── response.go                # Standardized response utilities
├── config.yaml                        # Configuration file
├── go.mod                             # Go module file
└── README.md                          # This file
```

## 🛠️ Setup Instructions

### Prerequisites

1. **Go** (version 1.19 or higher)
2. **PostgreSQL** (version 12 or higher)
3. **Git**

### Installation

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd clarityfin-api
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Configure the database**:
   - Create a PostgreSQL database named `clarityfin`
   - Update the `config.yaml` file with your database credentials:
   ```yaml
   database:
     dsn: "host=localhost user=postgres password=yourpassword dbname=clarityfin port=5432 sslmode=disable"
   ```

4. **Run the application**:
   ```bash
   go run cmd/api/main.go
   ```

The server will start on `http://localhost:8080`

## 📡 API Endpoints

### Authentication Endpoints

#### Register User
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "phone_number": "+1234567890",
  "password": "securepassword"
}
```

**Response**:
```json
{
  "success": true,
  "message": "Registration successful",
  "data": null
}
```

#### Login User
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "phone_number": "+1234567890",
  "password": "securepassword"
}
```

**Response**:
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### Protected Endpoints

#### Get Subscriptions
```http
GET /api/v1/subscriptions
Authorization: Bearer <your-jwt-token>
```

**Response**:
```json
{
  "success": true,
  "message": "Subscriptions retrieved successfully",
  "data": {
    "subscriptions": [
      {
        "id": 1,
        "name": "Netflix",
        "amount": 199,
        "user_id": 1,
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      }
    ],
    "total": 1
  }
}
```

#### Create Subscription
```http
POST /api/v1/subscriptions
Authorization: Bearer <your-jwt-token>
Content-Type: application/json

{
  "name": "Spotify",
  "amount": 119
}
```

#### Get Subscription by ID
```http
GET /api/v1/subscriptions/1
Authorization: Bearer <your-jwt-token>
```

## 🔧 Configuration

The application uses `config.yaml` for configuration:

```yaml
server:
  port: "8080"

database:
  dsn: "host=localhost user=postgres password=yourpassword dbname=clarityfin port=5432 sslmode=disable"

jwt:
  secret: "a-very-secret-key-that-is-long-and-secure"
```

## 🧪 Testing the API

### Using curl

1. **Register a new user**:
   ```bash
   curl -X POST http://localhost:8080/api/v1/auth/register \
     -H "Content-Type: application/json" \
     -d '{"phone_number": "+1234567890", "password": "testpassword"}'
   ```

2. **Login to get a token**:
   ```bash
   curl -X POST http://localhost:8080/api/v1/auth/login \
     -H "Content-Type: application/json" \
     -d '{"phone_number": "+1234567890", "password": "testpassword"}'
   ```

3. **Create a subscription**:
   ```bash
   curl -X POST http://localhost:8080/api/v1/subscriptions \
     -H "Authorization: Bearer <your-jwt-token>" \
     -H "Content-Type: application/json" \
     -d '{"name": "Netflix", "amount": 199}'
   ```

4. **Get subscriptions**:
   ```bash
   curl -X GET http://localhost:8080/api/v1/subscriptions \
     -H "Authorization: Bearer <your-jwt-token>"
   ```

## 🔒 Security Features

- **Password Hashing**: All passwords are hashed using bcrypt
- **JWT Authentication**: Secure token-based authentication
- **Input Validation**: Request validation using Gin's binding
- **Database Security**: Prepared statements via GORM
- **CORS Protection**: Cross-origin request handling

## 🏛️ Architecture Benefits

- **Maintainability**: Clear separation of concerns makes code easy to maintain
- **Testability**: Each layer can be tested independently
- **Scalability**: Easy to add new features and modify existing ones
- **Dependency Inversion**: High-level modules don't depend on low-level modules
- **Single Responsibility**: Each component has a single, well-defined purpose

## 🚀 Next Steps

This is the foundational backend for ClarityFin with clean architecture. Future enhancements could include:

- Unit and integration tests
- User profile management
- Real subscription data integration
- Payment processing
- Analytics and reporting
- Rate limiting
- Logging and monitoring
- API documentation with Swagger
- Docker containerization
- CI/CD pipeline

## 📝 License

This project is part of the ClarityFin application suite.
