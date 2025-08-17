package domain

import (
	"time"

	"gorm.io/gorm"
)

// Transaction represents the financial transaction domain entity
type Transaction struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	AccountID   uint           `json:"account_id" gorm:"not null"`
	Account     *Account       `json:"account,omitempty" gorm:"foreignKey:AccountID"`
	Type        string         `json:"type" gorm:"not null"` // credit, debit
	Amount      float64        `json:"amount" gorm:"not null"`
	Description string         `json:"description"`
	Category    string         `json:"category"`                          // food, transport, entertainment, etc.
	Status      string         `json:"status" gorm:"default:'completed'"` // pending, completed, failed
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// TransactionRepository defines the interface for transaction data operations
type TransactionRepository interface {
	Create(transaction *Transaction) error
	FindByID(id uint) (*Transaction, error)
	FindByAccountID(accountID uint) ([]*Transaction, error)
	FindByUserID(userID uint) ([]*Transaction, error)
	Update(transaction *Transaction) error
	Delete(id uint) error
}

// TransactionService defines the interface for transaction business logic
type TransactionService interface {
	CreateTransaction(accountID uint, transactionType, description, category string, amount float64) (*Transaction, error)
	GetAccountTransactions(accountID uint) ([]*Transaction, error)
	GetUserTransactions(userID uint) ([]*Transaction, error)
	GetTransactionByID(id uint) (*Transaction, error)
}

// TransactionUseCase defines the interface for transaction application logic
type TransactionUseCase interface {
	CreateTransaction(accountID uint, transactionType, description, category string, amount float64) (*Transaction, error)
	GetAccountTransactions(accountID uint) ([]*Transaction, error)
	GetUserTransactions(userID uint) ([]*Transaction, error)
	GetTransactionByID(id uint) (*Transaction, error)
}
