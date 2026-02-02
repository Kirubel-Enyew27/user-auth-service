package user

import (
	"net/http"
	"time"
	"user-auth-service/models"
	"user-auth-service/pkg/response"
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

func (usr *userService) Register(req models.RegisterUser) response.ErrorResponse {
	if err := req.Validate(); err != nil {
		usr.logger.Error("failed to validate user data", zap.Error(err))
		return response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "validation failed",
		}
	}

	user, _ := usr.userRepo.GetUserByNameOrPhone(req.Username, req.Phone)
	if user.ID != "" {
		if user.Username == req.Username {
			usr.logger.Error("username already exists", zap.String("username", user.Username))
			return response.ErrorResponse{
				StatusCode: http.StatusConflict,
				Message:    "username already exists",
			}
		}
		usr.logger.Error("phone already registered", zap.String("phone", user.Phone))
		return response.ErrorResponse{
			StatusCode: http.StatusConflict,
			Message:    "phone alreaady registered",
		}
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		usr.logger.Error("failed to hash password", zap.Error(err))
		return response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to hash password",
		}
	}

	user = models.User{
		ID:        uuid.NewString(),
		Username:  req.Username,
		Password:  hashedPassword,
		Phone:     req.Phone,
		Email:     req.Email,
		Role:      req.Role,
		Status:    string(models.StatusActive),
		CreatedAt: time.Now(),
	}

	errResp := usr.userRepo.Register(user)
	if errResp.Message != "" {
		return errResp
	}

	return response.ErrorResponse{}
}
