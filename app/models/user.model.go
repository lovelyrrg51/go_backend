package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id" gorm:"primary_key"`
	Username       string    `json:"username" gorm:"not null;unique"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email" gorm:"not null;unique"`
	Password       string    `json:"password"`
	IsDelete       bool      `json:"is_delete"`
	IsEmailConfirm bool      `json:"is_email_confirm`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
