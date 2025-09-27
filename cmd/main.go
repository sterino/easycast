package main

import (
	"database/sql"
	"easycart/internal/app"
	"easycart/internal/config"
	"easycart/migrations"
	"easycart/pkg/logger/sl"
	"fmt"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

//go:generate bin/generate.sh
func main() {
	if err := createRootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func createRootCmd() *cobra.Command {
	var application *app.App
	var err error
	var cfg *config.Config
	environment := os.Getenv("ENV")

	switch environment {
	case "local":
		cfg, err = config.NewLocal()
		if err != nil {
			slog.Error("Config create error", "error", err)
			os.Exit(1)
		}
		sl.Set(environment, cfg.Debug)
		slog.Info("Environment", "environment", environment)
		application, err = app.NewApp(cfg)
	default:
		cfg, err = config.New()
		if err != nil {
			slog.Error("Config create error", "error", err)
			os.Exit(1)
		}
		sl.Set(environment, cfg.Debug)
		application, err = app.NewApp(cfg)
	}
	if err != nil {
		slog.Error("Application create error", "error", err)
		os.Exit(1)
	}

	rootCmd := &cobra.Command{
		Use:   "run-server",
		Short: "Run server",
		Long:  "Run server is a default command",
		Run: func(cmd *cobra.Command, args []string) {

			go func() {
				if err := application.RunApp(); err != nil {
					slog.Warn(err.Error())
				}
			}()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit

			slog.Info("Graceful shutdown initiated...")
			if err := application.StopApp(); err != nil {
				slog.Error("Error during shutdown", "error", err)
			} else {
				slog.Info("Shutdown complete")
			}
		},
	}

	rootCmd.AddCommand(newMigrateCmd(cfg))

	return rootCmd
}

func newMigrateCmd(cfg *config.Config) *cobra.Command {
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Run migrations",
		Long:  "Run migrations. Specify action to run up/down migrations.",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Println("Run any of sub commands") },
	}

	db := func(cmd *cobra.Command) *sql.DB {
		db, err := migrations.NewSQLDB(&cfg.Database)
		if err != nil {
			cmd.Print(err)
			os.Exit(1)
		}
		return db
	}

	if err := migrations.Up(db(migrateCmd)); err != nil {
		slog.Error("Error migrations up", "error", err)
	}

	migrateCmd.AddCommand(
		&cobra.Command{
			Use:   "up",
			Short: "Run migrations up",
			Long:  "Run migrations UP to end.",
			Run: func(cmd *cobra.Command, args []string) {
				if err := migrations.Up(db(cmd)); err != nil {
					cmd.Print(err)
				}
			},
		},
		&cobra.Command{
			Use:   "down",
			Short: "Run migrations down",
			Long:  "Run migrations DOWN to the start.",
			Run: func(cmd *cobra.Command, args []string) {
				if err := migrations.Down(db(cmd)); err != nil {
					cmd.Print(err)
				}
			},
		},
		&cobra.Command{
			Use:   "steps N",
			Short: "Run migrations for given steps",
			Long:  "Run migrations {N} times. Steps can be positive (for UP migrations) or negative (for DOWN migrations)",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				steps, err := strconv.Atoi(args[0])
				if err != nil {
					cmd.Print(err)
					return
				}

				if err := migrations.Steps(db(cmd), steps); err != nil {
					cmd.Print(err)
				}
			},
		},
	)

	return migrateCmd
}
