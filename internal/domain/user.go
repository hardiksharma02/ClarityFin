package domain

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user domain entity
type User struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	PhoneNumber string         `json:"phone_number" gorm:"unique;not null"`
	Password    string         `json:"-" gorm:"not null"` // "-" means this field won't be serialized
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// UserRepository defines the interface for user data operations
type UserRepository interface {
	Create(user *User) error
	FindByPhoneNumber(phoneNumber string) (*User, error)
	FindByID(id uint) (*User, error)
	Update(user *User) error
	Delete(id uint) error
}

// UserService defines the interface for user business logic
type UserService interface {
	Register(phoneNumber, password string) error
	Authenticate(phoneNumber, password string) (*User, error)
	GetByID(id uint) (*User, error)
	GetByPhoneNumber(phoneNumber string) (*User, error)
}

// UserUseCase defines the interface for user application logic
type UserUseCase interface {
	Register(phoneNumber, password string) error
	Login(phoneNumber, password string) (string, error)
}
