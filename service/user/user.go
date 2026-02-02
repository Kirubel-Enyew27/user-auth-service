package user

import (
	"time"
	"user-auth-service/models"
	"user-auth-service/repo"
	"user-auth-service/service"
	"user-auth-service/utils"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type userService struct {
	userRepo repo.User
	logger   *zap.Logger
}

func NewService(userRepo repo.User, logger *zap.Logger) service.User {
	return &userService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (usr *userService) Register(req models.RegisterUser) error {
	if err := req.Validate(); err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := models.User{
		ID:        uuid.NewString(),
		Username:  req.Username,
		Password:  hashedPassword,
		Phone:     req.Phone,
		Email:     req.Email,
		Role:      req.Role,
		Status:    string(models.StatusActive),
		CreatedAt: time.Now(),
	}

	err = usr.userRepo.Register(user)
	if err != nil {
		return err
	}

	return nil
}
