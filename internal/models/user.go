package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"`
	Email        string        `gorm:"uniqueIndex;not null" json:"email"`
	Password     string        `gorm:"not null" json:"-"`
	FirstName    string        `gorm:"not null" json:"firstName"`
	LastName     string        `gorm:"not null" json:"lastName"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
	Role         string        `gorm:"type:varchar(20);not null;default:user" json:"role"`
}
