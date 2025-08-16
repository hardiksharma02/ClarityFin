package service

import (
	"github.com/hardiksharma/clarityfin-api/internal/domain"
)

// subscriptionUseCase implements the SubscriptionUseCase interface
type subscriptionUseCase struct {
	subscriptionService domain.SubscriptionService
}

// NewSubscriptionUseCase creates a new instance of SubscriptionUseCase
func NewSubscriptionUseCase(subscriptionService domain.SubscriptionService) domain.SubscriptionUseCase {
	return &subscriptionUseCase{
		subscriptionService: subscriptionService,
	}
}

// CreateSubscription handles subscription creation
func (uc *subscriptionUseCase) CreateSubscription(userID uint, name string, amount float64) (*domain.Subscription, error) {
	return uc.subscriptionService.CreateSubscription(userID, name, amount)
}

// GetUserSubscriptions handles retrieving user subscriptions
func (uc *subscriptionUseCase) GetUserSubscriptions(userID uint) ([]*domain.Subscription, error) {
	return uc.subscriptionService.GetUserSubscriptions(userID)
}

// GetSubscriptionByID handles retrieving a subscription by ID
func (uc *subscriptionUseCase) GetSubscriptionByID(id uint) (*domain.Subscription, error) {
	return uc.subscriptionService.GetSubscriptionByID(id)
}

// UpdateSubscription handles subscription updates
func (uc *subscriptionUseCase) UpdateSubscription(id uint, name string, amount float64) (*domain.Subscription, error) {
	return uc.subscriptionService.UpdateSubscription(id, name, amount)
}

// DeleteSubscription handles subscription deletion
func (uc *subscriptionUseCase) DeleteSubscription(id uint) error {
	return uc.subscriptionService.DeleteSubscription(id)
}
