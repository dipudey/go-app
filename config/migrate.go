package config

import (
	"log"

	"github.com/dipudey/go-app/internal/user"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&user.User{},
	)

	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("Database migrations completed")
}
