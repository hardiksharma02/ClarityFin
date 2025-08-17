package database

import (
	"fmt"
	"log"
	"strings"

	"github.com/hardiksharma/clarityfin-api/internal/config"
	"github.com/hardiksharma/clarityfin-api/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes the database connection.
func Connect(config config.DatabaseConfig) {
	var err error

	// Determine database type based on DSN
	if strings.Contains(config.DSN, "postgres") || strings.Contains(config.DSN, "localhost") {
		// PostgreSQL connection
		DB, err = gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	} else {
		// SQLite connection
		DB, err = gorm.Open(sqlite.Open(config.DSN), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connection successfully opened")

	// AutoMigrate will create the tables based on your GORM models
	err = DB.AutoMigrate(&domain.User{}, &domain.Subscription{}, &domain.OTP{}, &domain.Account{}, &domain.Transaction{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database migrated")
}
