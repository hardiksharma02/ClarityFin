package database

import (
	"fmt"
	"log"

	"github.com/hardiksharma/clarityfin-api/internal/config"
	"github.com/hardiksharma/clarityfin-api/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes the database connection.
func Connect(config config.DatabaseConfig) {
	var err error
	DB, err = gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connection successfully opened")

	// AutoMigrate will create the tables based on your GORM models
	err = DB.AutoMigrate(&domain.User{}, &domain.Subscription{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database migrated")
}
