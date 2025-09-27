package user

import (
	"github.com/google/uuid"
	"time"
)

const (
	UserTableName = "users"
)

type User struct {
	ID            uuid.UUID `db:"id"`
	Name          string    `db:"name"`
	Avatar        string    `db:"avatar"`
	PhoneNumber   string    `db:"phone_number"`
	PhoneVerified bool      `db:"phone_verified"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
	Password      string    `db:"password"`
}
