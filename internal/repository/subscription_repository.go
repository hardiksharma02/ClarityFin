package repository

import (
	"github.com/hardiksharma/clarityfin-api/internal/domain"
	"gorm.io/gorm"
)

// subscriptionRepository implements the SubscriptionRepository interface
type subscriptionRepository struct {
	db *gorm.DB
}

// NewSubscriptionRepository creates a new instance of SubscriptionRepository
func NewSubscriptionRepository(db *gorm.DB) domain.SubscriptionRepository {
	return &subscriptionRepository{db: db}
}

// Create creates a new subscription in the database
func (r *subscriptionRepository) Create(subscription *domain.Subscription) error {
	return r.db.Create(subscription).Error
}

// FindByID finds a subscription by ID
func (r *subscriptionRepository) FindByID(id uint) (*domain.Subscription, error) {
	var subscription domain.Subscription
	err := r.db.First(&subscription, id).Error
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}

// FindByUserID finds all subscriptions for a specific user
func (r *subscriptionRepository) FindByUserID(userID uint) ([]*domain.Subscription, error) {
	var subscriptions []*domain.Subscription
	err := r.db.Where("user_id = ?", userID).Find(&subscriptions).Error
	if err != nil {
		return nil, err
	}
	return subscriptions, nil
}

// Update updates an existing subscription
func (r *subscriptionRepository) Update(subscription *domain.Subscription) error {
	return r.db.Save(subscription).Error
}

// Delete deletes a subscription by ID
func (r *subscriptionRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Subscription{}, id).Error
}
