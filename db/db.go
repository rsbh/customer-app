package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/rsbh/customer-app/config"
)

func Connect(config *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", config.DATABASE_URL)
	return db, err
}
