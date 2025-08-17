package domain

import (
	"time"

	"gorm.io/gorm"
)

// Account represents the financial account domain entity
type Account struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id" gorm:"not null"`
	User        *User          `json:"user,omitempty" gorm:"foreignKey:UserID"`
	AccountType string         `json:"account_type" gorm:"not null"` // savings, checking, credit
	Balance     float64        `json:"balance" gorm:"default:0"`
	Currency    string         `json:"currency" gorm:"default:'USD'"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// AccountRepository defines the interface for account data operations
type AccountRepository interface {
	Create(account *Account) error
	FindByID(id uint) (*Account, error)
	FindByUserID(userID uint) ([]*Account, error)
	Update(account *Account) error
	Delete(id uint) error
}

// AccountService defines the interface for account business logic
type AccountService interface {
	CreateAccount(userID uint, accountType, currency string) (*Account, error)
	GetUserAccounts(userID uint) ([]*Account, error)
	GetAccountByID(id uint) (*Account, error)
	UpdateBalance(accountID uint, amount float64) error
}

// AccountUseCase defines the interface for account application logic
type AccountUseCase interface {
	CreateAccount(userID uint, accountType, currency string) (*Account, error)
	GetUserAccounts(userID uint) ([]*Account, error)
	GetAccountByID(id uint) (*Account, error)
}
