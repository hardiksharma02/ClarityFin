package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/hardiksharma/clarityfin-api/internal/config"
	"github.com/hardiksharma/clarityfin-api/internal/database"
	"github.com/hardiksharma/clarityfin-api/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware is a JWT middleware for protecting routes
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
			c.Abort()
			return
		}

		// This key should come from your config in a real app
		jwtKey := []byte("a-very-secret-key-that-is-long-and-secure")

		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user", claims.Subject)
		c.Next()
	}
}

func main() {
	// 1. Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// 2. Connect to the database
	database.Connect(cfg.Database)

	// 3. Set up the Gin router
	router := gin.Default()

	// Group API routes
	api := router.Group("/api/v1")
	{
		// Auth routes are public
		auth := api.Group("/auth")
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)

		// Subscription routes are protected
		subs := api.Group("/subscriptions")
		subs.Use(AuthMiddleware()) // Apply the middleware here
		subs.GET("/", handlers.GetSubscriptions)
	}

	// 4. Start the server
	log.Printf("Starting server on port %s", cfg.Server.Port)
	router.Run(":" + cfg.Server.Port)
}
