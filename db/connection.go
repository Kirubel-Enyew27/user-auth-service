package db

import (
	"database/sql"
	"fmt"
	"user-auth-service/config"
)

var DB *sql.DB

func Connect(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.POSTGRES_HOST, cfg.POSTGRES_PORT, cfg.POSTGRES_USERNAME, cfg.POSTGRES_PASSWORD, cfg.DB_NAME)

	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		return DB, fmt.Errorf("failed to open DB connection: %v", err)
	}

	if err := DB.Ping(); err != nil {
		return DB, fmt.Errorf("failed to ping DB: %v", err)
	}

	return DB, nil
}
