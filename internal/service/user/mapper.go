package user

import (
	udto "easycart/internal/controller/http/user/dto"
	"easycart/internal/domain"
)

func UserDtoToDomain(v udto.User, password string) *domain.User {
	return &domain.User{
		ID:            v.ID,
		Name:          v.Name,
		PhoneVerified: v.PhoneVerified,
		CreatedAt:     v.CreatedAt,
		UpdatedAt:     v.UpdatedAt,
		PhoneNumber:   v.PhoneNumber,
		Password:      password,
	}
}

func UserDomainToDto(v domain.User) udto.User {
	return udto.User{
		ID:            v.ID,
		Name:          v.Name,
		PhoneVerified: v.PhoneVerified,
		CreatedAt:     v.CreatedAt,
		UpdatedAt:     v.UpdatedAt,
		PhoneNumber:   v.PhoneNumber,
	}
}
