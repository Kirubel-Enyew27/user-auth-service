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
	"golang.org/x/crypto/bcrypt"
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

func (usr *userService) Login(req models.LoginRequest) (string, response.ErrorResponse) {
	if err := req.Validate(); err != nil {
		return "", response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid user input: " + err.Error(),
		}
	}

	user, errResp := usr.userRepo.GetUserByNameOrPhone(req.Username, "")
	if errResp.Message != "" || user.Username == "" {
		return "", errResp
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to generate token: " + err.Error(),
		}
	}

	return token, response.ErrorResponse{}
}

func (usr *userService) ChangePassword(req models.ChangePassword) response.ErrorResponse {
	if err := req.Validate(); err != nil {
		usr.logger.Error("validation failed", zap.Error(err))
		return response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid user input: " + err.Error(),
		}
	}

	user, errResp := usr.userRepo.GetUserByID(req.ID)
	if errResp.Message != "" || user.Username == "" {
		return errResp
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword))
	if err != nil {
		usr.logger.Error("the password doesn't match", zap.Error(err))
		return response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "the password doesn't match",
		}
	}

	newHashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		usr.logger.Error("failed to hash password", zap.Error(err))
		return response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	errResp = usr.userRepo.UpdatePassword(req.ID, newHashedPassword)
	if errResp.Message != "" {
		return errResp
	}

	return response.ErrorResponse{}

}

func (usr *userService) SuspendUser(id string) response.ErrorResponse {
	user, errResp := usr.userRepo.GetUserByID(id)
	if errResp.Message != "" || user.Username == "" {
		return errResp
	}

	errResp = usr.userRepo.SuspendUser(id)
	if errResp.Message != "" {
		return errResp
	}

	return response.ErrorResponse{}
}
