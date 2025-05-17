package database

import (
	"fmt"
	"os"

	"github.com/DevAthhh/auth-service/internal/infrastructure/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}

func SyncDB(db *gorm.DB) error {
	return db.AutoMigrate(&entity.User{})
}
