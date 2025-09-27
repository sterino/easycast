package dto

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Avatar        string    `json:"avatar"`
	PhoneNumber   string    `json:"phone_number"`
	PhoneVerified bool      `json:"phone_verified"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
