package user

import (
	"easycart/internal/config"
	"easycart/internal/service"
	"easycart/pkg/api/response"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	cfg *config.Config
	as  service.UserServicer
}

func NewAuthController(cfg *config.Config, as service.UserServicer) *UserController {
	return &UserController{
		cfg: cfg,
		as:  as,
	}
}

func (s *UserController) GetById(c echo.Context) error {
	id := c.Param("id")

	res, err := s.as.GetById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ValidationError(errors.New("get user by id error")))
	}

	return c.JSON(http.StatusOK, res)
}
