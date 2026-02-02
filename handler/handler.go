package handler

import (
	"user-auth-service/service"

	"go.uber.org/zap"
)

type userHandler struct {
	userService service.User
	logger      *zap.Logger
}

type User interface{}

func NewHandler(userService service.User, logger *zap.Logger) User {
	return &userHandler{
		userService: userService,
		logger:      logger,
	}
}
