# ClarityFin Backend API

A robust Go backend for the ClarityFin financial management application, built with Gin, GORM, and PostgreSQL.

## 🚀 Features

- **User Authentication**: JWT-based authentication with phone number and password
- **Database Integration**: PostgreSQL with GORM ORM
- **Configuration Management**: YAML-based configuration with Viper
- **RESTful API**: Clean API endpoints with proper HTTP status codes
- **Security**: Password hashing with bcrypt and JWT token validation

## 📁 Project Structure

```
clarityfin-api/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── database/
│   │   └── database.go          # Database connection and setup
│   ├── handlers/
│   │   ├── auth.go              # Authentication handlers
│   │   └── subscription.go      # Subscription handlers
│   └── models/
│       ├── user.go              # User model
│       └── subscription.go      # Subscription model
├── config.yaml                  # Configuration file
├── go.mod                       # Go module file
└── README.md                    # This file
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
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
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
  "subscriptions": [
    {
      "name": "Netflix",
      "amount": 199
    },
    {
      "name": "Spotify", 
      "amount": 119
    }
  ]
}
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

3. **Access protected endpoint**:
   ```bash
   curl -X GET http://localhost:8080/api/v1/subscriptions \
     -H "Authorization: Bearer <your-jwt-token>"
   ```

## 🔒 Security Features

- **Password Hashing**: All passwords are hashed using bcrypt
- **JWT Authentication**: Secure token-based authentication
- **Input Validation**: Request validation using Gin's binding
- **Database Security**: Prepared statements via GORM

## 🚀 Next Steps

This is the foundational backend for ClarityFin. Future enhancements could include:

- User profile management
- Real subscription data integration
- Payment processing
- Analytics and reporting
- Rate limiting
- Logging and monitoring
- Unit and integration tests

## 📝 License

This project is part of the ClarityFin application suite.
