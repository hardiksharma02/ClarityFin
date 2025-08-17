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
- **Mobile OTP Verification**: SMS-based OTP verification using Twilio or MSG91
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

#### Register User with OTP
```http
POST /api/v1/auth/register/otp
Content-Type: application/json

{
  "phone_number": "+1234567890",
  "password": "securepassword",
  "otp_code": "123456"
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

#### Send OTP
```http
POST /api/v1/otp/send
Content-Type: application/json

{
  "phone_number": "+1234567890"
}
```

**Response**:
```json
{
  "success": true,
  "message": "OTP sent successfully"
}
```

#### Verify OTP
```http
POST /api/v1/otp/verify
Content-Type: application/json

{
  "phone_number": "+1234567890",
  "code": "123456"
}
```

**Response**:
```json
{
  "success": true,
  "message": "OTP verified successfully"
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

sms:
  provider: "twilio"  # twilio or msg91
  twilio:
    account_sid: "your_twilio_account_sid"
    auth_token: "your_twilio_auth_token"
    from_number: "+1234567890"
  msg91:
    api_key: "your_msg91_api_key"
    sender_id: "CLARITY"
```

## 🧪 Testing the API

### Quick Start Testing

1. **Start the server**:
   ```bash
   go run cmd/api/main.go
   ```

2. **Open a new terminal and run the complete test suite**:
   ```bash
   # Test user registration
   curl -X POST http://localhost:8080/api/v1/auth/register \
     -H "Content-Type: application/json" \
     -d '{"phone_number": "+1234567890", "password": "testpassword"}'
   
   # Test user login and get JWT token
   TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
     -H "Content-Type: application/json" \
     -d '{"phone_number": "+1234567890", "password": "testpassword"}' \
     | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
   
   echo "JWT Token: $TOKEN"
   
   # Test getting subscriptions (should be empty initially)
   curl -X GET http://localhost:8080/api/v1/subscriptions/ \
     -H "Authorization: Bearer $TOKEN"
   
   # Test creating a subscription
   curl -X POST http://localhost:8080/api/v1/subscriptions/ \
     -H "Authorization: Bearer $TOKEN" \
     -H "Content-Type: application/json" \
     -d '{"name": "Netflix", "amount": 199}'
   
   # Test getting all subscriptions (should now show the created subscription)
   curl -X GET http://localhost:8080/api/v1/subscriptions/ \
     -H "Authorization: Bearer $TOKEN"
   ```

### Complete Test Suite

Here's a comprehensive test script you can run to test all endpoints:

```bash
#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}🚀 Starting ClarityFin API Test Suite${NC}"
echo "=================================="

# Base URL
BASE_URL="http://localhost:8080/api/v1"

# Test 1: User Registration
echo -e "\n${YELLOW}1. Testing User Registration${NC}"
echo "--------------------------------"
REGISTER_RESPONSE=$(curl -s -X POST $BASE_URL/auth/register \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+1234567890", "password": "testpassword"}')

echo "Response: $REGISTER_RESPONSE"

# Test 2: Duplicate Registration (should fail)
echo -e "\n${YELLOW}2. Testing Duplicate Registration${NC}"
echo "--------------------------------"
DUPLICATE_RESPONSE=$(curl -s -X POST $BASE_URL/auth/register \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+1234567890", "password": "testpassword"}')

echo "Response: $DUPLICATE_RESPONSE"

# Test 3: User Login
echo -e "\n${YELLOW}3. Testing User Login${NC}"
echo "--------------------------------"
LOGIN_RESPONSE=$(curl -s -X POST $BASE_URL/auth/login \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+1234567890", "password": "testpassword"}')

echo "Response: $LOGIN_RESPONSE"

# Extract JWT token
TOKEN=$(echo $LOGIN_RESPONSE | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
echo -e "${GREEN}JWT Token extracted: ${TOKEN:0:50}...${NC}"

# Test 4: Invalid Login
echo -e "\n${YELLOW}4. Testing Invalid Login${NC}"
echo "--------------------------------"
INVALID_LOGIN_RESPONSE=$(curl -s -X POST $BASE_URL/auth/login \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+1234567890", "password": "wrongpassword"}')

echo "Response: $INVALID_LOGIN_RESPONSE"

# Test 5: Get Subscriptions (Unauthorized)
echo -e "\n${YELLOW}5. Testing Unauthorized Access${NC}"
echo "--------------------------------"
UNAUTHORIZED_RESPONSE=$(curl -s -X GET $BASE_URL/subscriptions/)
echo "Response: $UNAUTHORIZED_RESPONSE"

# Test 6: Get Subscriptions (Authorized)
echo -e "\n${YELLOW}6. Testing Get Subscriptions (Empty)${NC}"
echo "--------------------------------"
GET_SUBS_RESPONSE=$(curl -s -X GET $BASE_URL/subscriptions/ \
  -H "Authorization: Bearer $TOKEN")
echo "Response: $GET_SUBS_RESPONSE"

# Test 7: Create Subscription
echo -e "\n${YELLOW}7. Testing Create Subscription${NC}"
echo "--------------------------------"
CREATE_SUB_RESPONSE=$(curl -s -X POST $BASE_URL/subscriptions/ \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "Netflix", "amount": 199}')

echo "Response: $CREATE_SUB_RESPONSE"

# Extract subscription ID
SUB_ID=$(echo $CREATE_SUB_RESPONSE | grep -o '"id":[0-9]*' | cut -d':' -f2)
echo -e "${GREEN}Subscription ID: $SUB_ID${NC}"

# Test 8: Get Subscription by ID
echo -e "\n${YELLOW}8. Testing Get Subscription by ID${NC}"
echo "--------------------------------"
GET_SUB_BY_ID_RESPONSE=$(curl -s -X GET $BASE_URL/subscriptions/$SUB_ID \
  -H "Authorization: Bearer $TOKEN")
echo "Response: $GET_SUB_BY_ID_RESPONSE"

# Test 9: Get All Subscriptions (should now have data)
echo -e "\n${YELLOW}9. Testing Get All Subscriptions (with data)${NC}"
echo "--------------------------------"
GET_ALL_SUBS_RESPONSE=$(curl -s -X GET $BASE_URL/subscriptions/ \
  -H "Authorization: Bearer $TOKEN")
echo "Response: $GET_ALL_SUBS_RESPONSE"

# Test 10: Create Another Subscription
echo -e "\n${YELLOW}10. Testing Create Another Subscription${NC}"
echo "--------------------------------"
CREATE_SUB2_RESPONSE=$(curl -s -X POST $BASE_URL/subscriptions/ \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "Spotify", "amount": 119}')

echo "Response: $CREATE_SUB2_RESPONSE"

# Test 11: Invalid Token
echo -e "\n${YELLOW}11. Testing Invalid Token${NC}"
echo "--------------------------------"
INVALID_TOKEN_RESPONSE=$(curl -s -X GET $BASE_URL/subscriptions/ \
  -H "Authorization: Bearer invalid_token")
echo "Response: $INVALID_TOKEN_RESPONSE"

# Test 12: Invalid Input Validation
echo -e "\n${YELLOW}12. Testing Input Validation${NC}"
echo "--------------------------------"
INVALID_INPUT_RESPONSE=$(curl -s -X POST $BASE_URL/subscriptions/ \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "", "amount": -100}')

echo "Response: $INVALID_INPUT_RESPONSE"

echo -e "\n${GREEN}✅ Test Suite Completed!${NC}"
echo "=================================="
```

### Manual Testing Steps

#### 1. **Start the Server**
```bash
go run cmd/api/main.go
```

#### 2. **Test User Registration**
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+1234567890", "password": "testpassword"}'
```

**Expected Response:**
```json
{
  "success": true,
  "message": "Registration successful"
}
```

#### 3. **Test User Login**
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+1234567890", "password": "testpassword"}'
```

**Expected Response:**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### 4. **Test Protected Endpoints**
```bash
# Get JWT token first
TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+1234567890", "password": "testpassword"}' \
  | grep -o '"token":"[^"]*"' | cut -d'"' -f4)

# Get subscriptions
curl -X GET http://localhost:8080/api/v1/subscriptions/ \
  -H "Authorization: Bearer $TOKEN"

# Create subscription
curl -X POST http://localhost:8080/api/v1/subscriptions/ \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "Netflix", "amount": 199}'
```

### Testing with Different Tools

#### **Using Postman**
1. Import the following collection:
```json
{
  "info": {
    "name": "ClarityFin API",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "Register User",
      "request": {
        "method": "POST",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n  \"phone_number\": \"+1234567890\",\n  \"password\": \"testpassword\"\n}"
        },
        "url": "http://localhost:8080/api/v1/auth/register"
      }
    },
    {
      "name": "Login User",
      "request": {
        "method": "POST",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n  \"phone_number\": \"+1234567890\",\n  \"password\": \"testpassword\"\n}"
        },
        "url": "http://localhost:8080/api/v1/auth/login"
      }
    },
    {
      "name": "Get Subscriptions",
      "request": {
        "method": "GET",
        "header": [
          {
            "key": "Authorization",
            "value": "Bearer {{jwt_token}}"
          }
        ],
        "url": "http://localhost:8080/api/v1/subscriptions/"
      }
    }
  ]
}
```

#### **Using Insomnia**
1. Create a new request collection
2. Add the endpoints with the same structure as above
3. Use environment variables for the JWT token

### Database Testing

#### **Check SQLite Database**
```bash
# View the database file
ls -la *.db

# Use SQLite CLI to inspect data (if installed)
sqlite3 clarityfin.db

# Inside SQLite CLI:
.tables                    # Show all tables
SELECT * FROM users;       # View users
SELECT * FROM subscriptions; # View subscriptions
.quit                      # Exit SQLite
```

### Error Testing

#### **Test Invalid Scenarios**
```bash
# Test duplicate registration
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+1234567890", "password": "testpassword"}'

# Test invalid login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+1234567890", "password": "wrongpassword"}'

# Test unauthorized access
curl -X GET http://localhost:8080/api/v1/subscriptions/

# Test invalid token
curl -X GET http://localhost:8080/api/v1/subscriptions/ \
  -H "Authorization: Bearer invalid_token"

# Test invalid input
curl -X POST http://localhost:8080/api/v1/subscriptions/ \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "", "amount": -100}'
```

### Performance Testing

#### **Load Testing with Apache Bench**
```bash
# Install Apache Bench (if not available)
# macOS: brew install httpd
# Ubuntu: sudo apt-get install apache2-utils

# Test registration endpoint
ab -n 100 -c 10 -p register_data.json -T application/json http://localhost:8080/api/v1/auth/register

# Test login endpoint
ab -n 100 -c 10 -p login_data.json -T application/json http://localhost:8080/api/v1/auth/login
```

### Automated Testing

#### **Create a Test Script**
Save the complete test suite above as `test_api.sh` and run:
```bash
chmod +x test_api.sh
./test_api.sh
```

This comprehensive testing guide ensures you can thoroughly test all aspects of the ClarityFin API!

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
