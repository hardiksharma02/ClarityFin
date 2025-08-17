package dto

// SendOTPRequest represents the request body for sending OTP
type SendOTPRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required" validate:"required,min=10,max=15"`
}

// VerifyOTPRequest represents the request body for verifying OTP
type VerifyOTPRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required" validate:"required"`
	Code        string `json:"code" binding:"required" validate:"required,len=6"`
}

// OTPResponse represents the response for OTP operations
type OTPResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}
