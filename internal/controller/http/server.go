package http

import (
	"fmt"
	"net/http"
	"time"

	//_ "easycart/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NewServer(cfg *Config, authRouter http.Handler, userRouter http.Handler) *http.Server {
	mux := http.NewServeMux()

	router := echo.New()
	router.GET("/swagger/*", echoSwagger.WrapHandler)
	router.OPTIONS("/*any", func(c echo.Context) error { return c.NoContent(http.StatusOK) })

	mux.Handle("/", http.StripPrefix("/api", CORS(router)))
	mux.Handle("/swagger/", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))
	mux.Handle("/health-check/", CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })))
	mux.Handle("/api/auth", http.StripPrefix("/api/auth", CORS(authRouter)))
	mux.Handle("/api/user", http.StripPrefix("/api/user", CORS(userRouter)))

	return &http.Server{
		Addr:              fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		Handler:           mux,
		ReadHeaderTimeout: time.Second * 5,
	}
}
