package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hardiksharma/clarityfin-api/internal/config"
	"github.com/hardiksharma/clarityfin-api/internal/database"
	"github.com/hardiksharma/clarityfin-api/internal/handlers"
	"github.com/hardiksharma/clarityfin-api/internal/middleware"
	"github.com/hardiksharma/clarityfin-api/internal/repository"
	"github.com/hardiksharma/clarityfin-api/internal/service"
)

func main() {
	// 1. Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// 2. Connect to the database
	database.Connect(cfg.Database)

	// 3. Initialize repositories
	userRepo := repository.NewUserRepository(database.DB)
	subscriptionRepo := repository.NewSubscriptionRepository(database.DB)

	// 4. Initialize services
	userService := service.NewUserService(userRepo, cfg.JWT.Secret)
	subscriptionService := service.NewSubscriptionService(subscriptionRepo, userRepo)

	// 5. Initialize use cases
	userUseCase := service.NewUserUseCase(userService)
	subscriptionUseCase := service.NewSubscriptionUseCase(subscriptionService)

	// 6. Initialize handlers
	authHandler := handlers.NewAuthHandler(userUseCase)
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionUseCase, userService)

	// 7. Set up the Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Group API routes
	api := router.Group("/api/v1")
	{
		// Auth routes are public
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Subscription routes are protected
		subs := api.Group("/subscriptions")
		subs.Use(middleware.AuthMiddleware(cfg.JWT.Secret))
		{
			subs.GET("/", subscriptionHandler.GetSubscriptions)
			subs.POST("/", subscriptionHandler.CreateSubscription)
			subs.GET("/:id", subscriptionHandler.GetSubscriptionByID)
		}
	}

	// 8. Start the server
	log.Printf("Starting server on port %s", cfg.Server.Port)
	router.Run(":" + cfg.Server.Port)
}
