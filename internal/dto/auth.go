package dto

// RegisterRequest represents the request body for user registration
type RegisterRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required" validate:"required,min=10,max=15"`
	Password    string `json:"password" binding:"required" validate:"required,min=6"`
	OTPCode     string `json:"otp_code,omitempty" validate:"omitempty,len=6"`
}

// LoginRequest represents the request body for user login
type LoginRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required" validate:"required"`
	Password    string `json:"password" binding:"required" validate:"required"`
}

// AuthResponse represents the response for authentication operations
type AuthResponse struct {
	Token   string        `json:"token,omitempty"`
	User    *UserResponse `json:"user,omitempty"`
	Message string        `json:"message,omitempty"`
}

// UserResponse represents the user data in API responses
type UserResponse struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
}
