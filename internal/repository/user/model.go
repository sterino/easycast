package user

import (
	"github.com/google/uuid"
	"time"
)

const (
	userTableName = "users"
)

type CreateUser struct {
	ID            uuid.UUID `db:"id"`
	Name          string    `db:"name"`
	Avatar        string    `db:"avatar"`
	PhoneNumber   string    `db:"phone_number"`
	PhoneVerified bool      `db:"phone_verified"`
	Password      string    `db:"password"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
