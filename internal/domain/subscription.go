package domain

import (
	"time"

	"gorm.io/gorm"
)

// Subscription represents the subscription domain entity
type Subscription struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	Amount    float64        `json:"amount" gorm:"not null"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	User      *User          `json:"user,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// SubscriptionRepository defines the interface for subscription data operations
type SubscriptionRepository interface {
	Create(subscription *Subscription) error
	FindByID(id uint) (*Subscription, error)
	FindByUserID(userID uint) ([]*Subscription, error)
	Update(subscription *Subscription) error
	Delete(id uint) error
}

// SubscriptionService defines the interface for subscription business logic
type SubscriptionService interface {
	CreateSubscription(userID uint, name string, amount float64) (*Subscription, error)
	GetUserSubscriptions(userID uint) ([]*Subscription, error)
	GetSubscriptionByID(id uint) (*Subscription, error)
	UpdateSubscription(id uint, name string, amount float64) (*Subscription, error)
	DeleteSubscription(id uint) error
}

// SubscriptionUseCase defines the interface for subscription application logic
type SubscriptionUseCase interface {
	CreateSubscription(userID uint, name string, amount float64) (*Subscription, error)
	GetUserSubscriptions(userID uint) ([]*Subscription, error)
	GetSubscriptionByID(id uint) (*Subscription, error)
	UpdateSubscription(id uint, name string, amount float64) (*Subscription, error)
	DeleteSubscription(id uint) error
}
