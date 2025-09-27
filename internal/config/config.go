package config

import (
	"log"

	"easycart/internal/controller/http"
	"easycart/migrations"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Logger struct {
	Level string `env:"LOG_LEVEL" envDefault:"debug"`
}

type Config struct {
	Debug    bool `env:"DEBUG" envDefault:"false"`
	Logger   Logger
	Server   http.Config
	Database migrations.Config
}

// New creates new config from env.
func New() (*Config, error) {
	var conf Config
	if err := env.Parse(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func NewLocal() (*Config, error) {
	var conf Config
	if err := godotenv.Load(".local.env"); err != nil {
		log.Println("Warning: No .env file found")
	}
	if err := env.Parse(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
