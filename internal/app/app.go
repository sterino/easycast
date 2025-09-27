package app

import (
	"context"
	"easycart/internal/config"
	"easycart/migrations"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	controller "easycart/internal/controller/http"
	authCtrl "easycart/internal/controller/http/auth"
	userCtrl "easycart/internal/controller/http/user"
	userStorage "easycart/internal/repository/user"
	userService "easycart/internal/service/user"
)

type App struct {
	Server *http.Server
	Cfg    *config.Config
	Logger *slog.Logger
}

func NewApp(cfg *config.Config) (*App, error) {
	// logger
	logger := slog.Default()
	// database
	db, err := migrations.NewGoquDB(&cfg.Database)
	if err != nil {
		logger.Error("unable to connect database", "error", err)
	}
	userStorage := userStorage.NewUserStorage(db)

	//services
	userService := userService.NewUserService(userStorage)

	// controllers
	userController := userCtrl.NewAuthController(cfg, userService)
	authController := authCtrl.NewAuthController(cfg, userService)

	// routes
	userRoute := userCtrl.NewRoutes(userController)
	authRoute := authCtrl.NewRoutes(authController)
	// server
	httpServer := controller.NewServer(&cfg.Server, authRoute, userRoute)

	return &App{Server: httpServer, Cfg: cfg, Logger: logger}, nil
}

func (a *App) RunApp() error {
	a.Logger.Info("starting server...", "address", fmt.Sprintf("http://localhost:%d", a.Cfg.Server.Port))
	if err := a.Server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (a *App) StopApp() error {
	a.Logger.Info("stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		a.Logger.Error("unable to shutdown server", "error", err)
		return err
	}
	a.Logger.Info("server stopped")

	return nil
}
