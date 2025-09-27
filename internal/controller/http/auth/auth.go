package auth

import (
	"easycart/internal/config"
	"easycart/internal/service"
)

type AuthController struct {
	cfg *config.Config
	as service.
}

func NewAuthController(cfg *config.Config, as service.) *AuthController {
	return &AuthController{
		cfg: cfg,
		as: as,
	}
}