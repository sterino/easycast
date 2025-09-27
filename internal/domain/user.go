package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID            uuid.UUID
	Name          string
	Avatar        string
	PhoneNumber   string
	PhoneVerified bool

	Password string
}

type AuthUser struct {
	PhoneNumber string
	Password    string
}

type AuthToken struct {
	Token string
}

type CreateUser struct {
	ID            uuid.UUID
	Name          string
	Avatar        string
	PhoneNumber   string
	PhoneVerified bool
	Password      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
