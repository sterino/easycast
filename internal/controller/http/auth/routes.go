package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"

	controller "easycart/internal/controller/http"
)

func NewRoutes(h *AuthController) http.Handler {
	router := echo.New()
	v1 := router.Group("/v1")
	v1.POST("/auth/register", h.)
	v1.POST("/auth/login", h.)
	v1.POST("/auth/register/verify", h.)


	return router
}
