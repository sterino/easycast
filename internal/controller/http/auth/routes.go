package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewRoutes(con *AuthController) http.Handler {
	router := echo.New()
	router.POST("/auth/register", con.Register)
	router.POST("/auth/login", con.Login)
	//router.POST("/auth/register/verify", h.)

	return router
}
