package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string    `gorm:"unique"`
	UUID     uuid.UUID `gorm:"unique"`
	Password string
}
