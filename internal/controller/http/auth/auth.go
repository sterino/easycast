package auth

import (
	"easycart/internal/config"
	"easycart/internal/controller/http/auth/dto"
	"easycart/internal/service"
	"easycart/pkg/api/response"
	"errors"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type AuthController struct {
	cfg *config.Config
	as  service.UserServicer
}

func NewAuthController(cfg *config.Config, as service.UserServicer) *AuthController {
	return &AuthController{
		cfg: cfg,
		as:  as,
	}
}

func (s *AuthController) Register(c echo.Context) error {
	req := dto.RegisterRequest{}
	if err := c.Bind(&req); err != nil {
		slog.Error("Bind error:", "err", err)
		return c.JSON(http.StatusBadRequest, response.ValidationError(errors.New("invalid request body")))
	}

	//if claimsFilters.IsAdmin != nil {
	//	if !*claimsFilters.IsAdmin {
	//		filial := c.Request().Context().Value(domain.CtxSpf).(string)
	//	}
	//}

	id, err := s.as.Register(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ValidationError(errors.New("register error")))
	}

	return c.JSON(http.StatusCreated, id)
}

func (s *AuthController) Login(c echo.Context) error {
	req := dto.LoginRequest{}
	if err := c.Bind(&req); err != nil {
		slog.Error("Bind error:", "err", err)
		return c.JSON(http.StatusBadRequest, response.ValidationError(errors.New("invalid request body")))
	}

	token, _, err := s.as.Login(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ValidationError(errors.New("login error")))
	}

	return c.JSON(http.StatusOK, token)
}
