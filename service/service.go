package service

import (
	"user-auth-service/repo"

	"go.uber.org/zap"
)

type userService struct {
	userRepo repo.User
	logger   *zap.Logger
}

type User interface {
}

func NewService(userRepo repo.User, logger *zap.Logger) User {
	return &userService{
		userRepo: userRepo,
		logger:   logger,
	}
}
