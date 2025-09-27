package migrations

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

// NewSQLDB returns raw sql.DB entity.
func NewSQLDB(config *Config) (*sql.DB, error) {
	if config == nil {
		return nil, errors.New("missing config")
	}
	dsn := config.DSN()
	if config.SSLMode == "" {
		dsn += " sslmode=disable"
	} else {
		dsn += fmt.Sprintf(" sslmode=%s", config.SSLMode)
	}
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	if config.MaxConns != 0 {
		sqlDB.SetMaxOpenConns(config.MaxConns)
	} else {
		sqlDB.SetMaxOpenConns(15)
	}

	// SetMaxIdleConns sets the maximum number of idle connections to the database.
	if config.MaxIdleConnections != 0 {
		sqlDB.SetMaxIdleConns(config.MaxIdleConnections)
	} else {
		sqlDB.SetMaxIdleConns(10)
	}

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	if config.MaxConnLifeTimeInSeconds != 0 {
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.MaxConnLifeTimeInSeconds))
	} else {
		sqlDB.SetConnMaxLifetime(25 * time.Minute)
	}

	// SetConnMaxIdletime sets the maximum amount of time an idle connection may be reused.
	if config.MaxConnIdleTimeInSeconds != 0 {
		sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(config.MaxConnIdleTimeInSeconds))
	} else {
		sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	}

	return sqlDB, nil
}

func Dial(ctx context.Context, url string) (*pgxpool.Pool, error) {
	conf, cfgErr := pgxpool.ParseConfig(url)

	if cfgErr != nil {
		return nil, cfgErr
	}
	conf.MaxConns = 20
	conf.MinConns = 10
	conf.MaxConnIdleTime = 10 * time.Second

	conn, connErr := pgxpool.ConnectConfig(ctx, conf)

	if connErr != nil {
		return nil, connErr
	}
	if pingErr := conn.Ping(ctx); pingErr != nil {
		return nil, pingErr
	}
	return conn, nil
}

func NewGoquDB(config *Config) (*goqu.Database, error) {
	sqlDB, err := NewSQLDB(config)
	if err != nil {
		return nil, err
	}

	dialect := goqu.Dialect("postgres")
	db := dialect.DB(sqlDB)

	return db, nil
}
