package domain

import (
	"time"
)

// OTP represents the OTP domain entity
type OTP struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	PhoneNumber string    `json:"phone_number" gorm:"not null"`
	Code        string    `json:"code" gorm:"not null"`
	ExpiresAt   time.Time `json:"expires_at" gorm:"not null"`
	IsUsed      bool      `json:"is_used" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// OTPRepository defines the interface for OTP data operations
type OTPRepository interface {
	Create(otp *OTP) error
	FindByPhoneNumberAndCode(phoneNumber, code string) (*OTP, error)
	MarkAsUsed(id uint) error
	DeleteExpired() error
}

// OTPService defines the interface for OTP business logic
type OTPService interface {
	GenerateOTP(phoneNumber string) error
	VerifyOTP(phoneNumber, code string) (bool, error)
	SendOTP(phoneNumber, code string) error
}

// OTPUseCase defines the interface for OTP application logic
type OTPUseCase interface {
	SendOTP(phoneNumber string) error
	VerifyOTP(phoneNumber, code string) (bool, error)
}
