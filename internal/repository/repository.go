package repository

import (
	"context"
	"easycart/internal/domain"
	"github.com/google/uuid"
)

type UserStorager interface {
	Create(ctx context.Context, user domain.CreateUser) (uuid.UUID, error)
	Find(ctx context.Context, id, phoneNumber string) ([]domain.User, error)
}
