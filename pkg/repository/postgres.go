package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	subscriptionTable = "Subscriptions"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s password=%s user=%s sslmode=%s", cfg.Host, cfg.Port, cfg.DBName, cfg.Password, cfg.Username, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
