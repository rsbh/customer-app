package db

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/rsbh/customer-app/config"
)

func MigrationsUp(config *config.Config) error {
	if config.DATABASE_URL == "" {
		return errors.New("DATABASE_URL is empty")
	}
	m, err := migrate.New(
		"file://db/migrations",
		config.DATABASE_URL)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func Connect(config *config.Config) (*sqlx.DB, error) {
	if config.DATABASE_URL == "" {
		return nil, errors.New("DATABASE_URL is empty")
	}
	db, err := sqlx.Connect("postgres", config.DATABASE_URL)
	return db, err
}
