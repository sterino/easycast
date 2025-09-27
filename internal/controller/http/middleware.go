package http

import (
	"easycart/pkg/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)
	})
}

const (
	ContextUserID = "user_id"
	ContextEmail  = "email"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Authorization header missing or invalid"})
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			// валидная ли структура токена (без расшифровки)
			if !jwt.Sanitize(token) {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token structure"})
			}

			secretKey := jwt.ProvideSecretKey()

			// расшифровка и проверка
			claims, err := jwt.Decode(token, secretKey)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
			}

			// Прокидываем данные в контекст
			c.Set(ContextUserID, claims.UUID)
			c.Set(ContextEmail, claims.PhoneNumber)

			return next(c)
		}
	}
}
