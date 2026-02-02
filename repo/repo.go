package repo

import (
	"database/sql"

	"go.uber.org/zap"
)

type userRepo struct {
	db     *sql.DB
	logger *zap.Logger
}

type User interface {
}

func NewRepo(db *sql.DB, logger *zap.Logger) User {
	return &userRepo{
		db:     db,
		logger: logger,
	}
}
