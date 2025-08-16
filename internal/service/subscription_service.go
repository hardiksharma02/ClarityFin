package service

import (
	"errors"

	"github.com/hardiksharma/clarityfin-api/internal/domain"
)

// subscriptionService implements the SubscriptionService interface
type subscriptionService struct {
	subscriptionRepo domain.SubscriptionRepository
	userRepo         domain.UserRepository
}

// NewSubscriptionService creates a new instance of SubscriptionService
func NewSubscriptionService(subscriptionRepo domain.SubscriptionRepository, userRepo domain.UserRepository) domain.SubscriptionService {
	return &subscriptionService{
		subscriptionRepo: subscriptionRepo,
		userRepo:         userRepo,
	}
}

// CreateSubscription creates a new subscription for a user
func (s *subscriptionService) CreateSubscription(userID uint, name string, amount float64) (*domain.Subscription, error) {
	// Verify user exists
	_, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	subscription := &domain.Subscription{
		Name:   name,
		Amount: amount,
		UserID: userID,
	}

	err = s.subscriptionRepo.Create(subscription)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}

// GetUserSubscriptions retrieves all subscriptions for a user
func (s *subscriptionService) GetUserSubscriptions(userID uint) ([]*domain.Subscription, error) {
	// Verify user exists
	_, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return s.subscriptionRepo.FindByUserID(userID)
}

// GetSubscriptionByID retrieves a subscription by ID
func (s *subscriptionService) GetSubscriptionByID(id uint) (*domain.Subscription, error) {
	return s.subscriptionRepo.FindByID(id)
}

// UpdateSubscription updates an existing subscription
func (s *subscriptionService) UpdateSubscription(id uint, name string, amount float64) (*domain.Subscription, error) {
	subscription, err := s.subscriptionRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("subscription not found")
	}

	subscription.Name = name
	subscription.Amount = amount

	err = s.subscriptionRepo.Update(subscription)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}

// DeleteSubscription deletes a subscription
func (s *subscriptionService) DeleteSubscription(id uint) error {
	// Verify subscription exists
	_, err := s.subscriptionRepo.FindByID(id)
	if err != nil {
		return errors.New("subscription not found")
	}

	return s.subscriptionRepo.Delete(id)
}
