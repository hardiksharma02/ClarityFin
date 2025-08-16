package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hardiksharma/clarityfin-api/internal/domain"
	"github.com/hardiksharma/clarityfin-api/internal/dto"
	"github.com/hardiksharma/clarityfin-api/pkg/response"
)

// SubscriptionHandler handles subscription-related HTTP requests
type SubscriptionHandler struct {
	subscriptionUseCase domain.SubscriptionUseCase
	userService         domain.UserService
}

// NewSubscriptionHandler creates a new instance of SubscriptionHandler
func NewSubscriptionHandler(subscriptionUseCase domain.SubscriptionUseCase, userService domain.UserService) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionUseCase: subscriptionUseCase,
		userService:         userService,
	}
}

// GetSubscriptions retrieves all subscriptions for the authenticated user
func (h *SubscriptionHandler) GetSubscriptions(c *gin.Context) {
	userPhone, exists := c.Get("user_phone")
	if !exists {
		response.Unauthorized(c, "User not authenticated")
		return
	}

	// Get user by phone number to get user ID
	user, err := h.userService.GetByPhoneNumber(userPhone.(string))
	if err != nil {
		response.InternalServerError(c, "Failed to get user")
		return
	}

	subscriptions, err := h.subscriptionUseCase.GetUserSubscriptions(user.ID)
	if err != nil {
		response.InternalServerError(c, "Failed to get subscriptions")
		return
	}

	// Convert domain models to DTOs
	subscriptionResponses := make([]*dto.SubscriptionResponse, len(subscriptions))
	for i, sub := range subscriptions {
		subscriptionResponses[i] = &dto.SubscriptionResponse{
			ID:        sub.ID,
			Name:      sub.Name,
			Amount:    sub.Amount,
			UserID:    sub.UserID,
			CreatedAt: sub.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt: sub.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}

	response.Success(c, dto.SubscriptionsResponse{
		Subscriptions: subscriptionResponses,
		Total:         len(subscriptionResponses),
	}, "Subscriptions retrieved successfully")
}

// CreateSubscription creates a new subscription
func (h *SubscriptionHandler) CreateSubscription(c *gin.Context) {
	var req dto.CreateSubscriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	userPhone, exists := c.Get("user_phone")
	if !exists {
		response.Unauthorized(c, "User not authenticated")
		return
	}

	// Get user by phone number to get user ID
	user, err := h.userService.GetByPhoneNumber(userPhone.(string))
	if err != nil {
		response.InternalServerError(c, "Failed to get user")
		return
	}

	subscription, err := h.subscriptionUseCase.CreateSubscription(user.ID, req.Name, req.Amount)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	subscriptionResponse := &dto.SubscriptionResponse{
		ID:        subscription.ID,
		Name:      subscription.Name,
		Amount:    subscription.Amount,
		UserID:    subscription.UserID,
		CreatedAt: subscription.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: subscription.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}

	response.Success(c, subscriptionResponse, "Subscription created successfully")
}

// GetSubscriptionByID retrieves a specific subscription
func (h *SubscriptionHandler) GetSubscriptionByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid subscription ID")
		return
	}

	subscription, err := h.subscriptionUseCase.GetSubscriptionByID(uint(id))
	if err != nil {
		response.NotFound(c, "Subscription not found")
		return
	}

	subscriptionResponse := &dto.SubscriptionResponse{
		ID:        subscription.ID,
		Name:      subscription.Name,
		Amount:    subscription.Amount,
		UserID:    subscription.UserID,
		CreatedAt: subscription.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: subscription.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}

	response.Success(c, subscriptionResponse, "Subscription retrieved successfully")
}
