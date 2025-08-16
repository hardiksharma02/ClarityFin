package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSubscriptions returns a list of mock subscriptions
func GetSubscriptions(c *gin.Context) {
	// In a real app, you would get the user ID from the context
	// (set by the middleware) and fetch their actual subscriptions from the DB.

	// For now, let's return some mock data.
	mockSubscriptions := []gin.H{
		{"name": "Netflix", "amount": 199},
		{"name": "Spotify", "amount": 119},
	}

	c.JSON(http.StatusOK, gin.H{"subscriptions": mockSubscriptions})
}
