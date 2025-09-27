package service

import (
	"context"
	adto "easycart/internal/controller/http/auth/dto"
	udto "easycart/internal/controller/http/user/dto"

	"github.com/google/uuid"
)

type UserServicer interface {
	GetById(context.Context, string) (*udto.User, error)
	Login(ctx context.Context, req adto.LoginRequest) (string, int64, error)
	List(ctx context.Context) (*[]udto.User, error)
	Register(ctx context.Context, req adto.RegisterRequest) (uuid.UUID, error)
}
