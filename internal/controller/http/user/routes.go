package user

import (
	controller "easycart/internal/controller/http"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewRoutes(con *UserController) http.Handler {
	router := echo.New()
	router.Use(controller.JWTMiddleware())
	router.POST("/user", con.GetById)
	//router.POST("/auth/register/verify", h.)

	return router
}
