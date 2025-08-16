package dto

// CreateSubscriptionRequest represents the request body for creating a subscription
type CreateSubscriptionRequest struct {
	Name   string  `json:"name" binding:"required" validate:"required,min=1,max=100"`
	Amount float64 `json:"amount" binding:"required" validate:"required,min=0"`
}

// UpdateSubscriptionRequest represents the request body for updating a subscription
type UpdateSubscriptionRequest struct {
	Name   string  `json:"name" validate:"required,min=1,max=100"`
	Amount float64 `json:"amount" validate:"required,min=0"`
}

// SubscriptionResponse represents the subscription data in API responses
type SubscriptionResponse struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Amount    float64 `json:"amount"`
	UserID    uint    `json:"user_id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// SubscriptionsResponse represents the response for subscription list operations
type SubscriptionsResponse struct {
	Subscriptions []*SubscriptionResponse `json:"subscriptions"`
	Total         int                     `json:"total"`
}
