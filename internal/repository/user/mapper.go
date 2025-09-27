package user

import "easycart/internal/domain"

func CreateUserToModel(v domain.CreateUser) CreateUser {
	return CreateUser{
		ID:            v.ID,
		Name:          v.Name,
		Avatar:        v.Avatar,
		Password:      v.Password,
		CreatedAt:     v.CreatedAt,
		UpdatedAt:     v.UpdatedAt,
		PhoneNumber:   v.PhoneNumber,
		PhoneVerified: v.PhoneVerified,
	}
}
