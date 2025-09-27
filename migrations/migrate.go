package migrations

import (
	"database/sql"
	"embed"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed data/*.sql
var fs embed.FS

func newMigrateInstanceWithSQLInstance(db *sql.DB) (*migrate.Migrate, error) {
	data, err := iofs.New(fs, "data")
	if err != nil {
		return nil, err
	}

	dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	return migrate.NewWithInstance("iofs", data, "postgres", dbDriver)
}

func Up(db *sql.DB) error {
	m, err := newMigrateInstanceWithSQLInstance(db)
	if err != nil {
		return err
	}
	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}
	return err
}

func Down(db *sql.DB) error {
	m, err := newMigrateInstanceWithSQLInstance(db)
	if err != nil {
		return err
	}
	err = m.Down()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}
	return err
}

func Steps(db *sql.DB, stepCount int) error {
	m, err := newMigrateInstanceWithSQLInstance(db)
	if err != nil {
		return err
	}
	err = m.Steps(stepCount)
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}
	return err
}
