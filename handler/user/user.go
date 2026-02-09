package user

import (
	"net/http"
	"user-auth-service/handler"
	"user-auth-service/models"
	"user-auth-service/pkg/response"
	"user-auth-service/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type userHandler struct {
	userService service.User
	logger      *zap.Logger
}

func NewHandler(userService service.User, logger *zap.Logger) handler.User {
	return &userHandler{
		userService: userService,
		logger:      logger,
	}
}

func (usr *userHandler) Register(c *gin.Context) {
	var req models.RegisterUser

	if err := c.ShouldBindJSON(&req); err != nil {
		usr.logger.Error("failed to bind request")
		response.SendErrorResponse(c, &response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	errResp := usr.userService.Register(req)
	if errResp.Message != "" {
		response.SendErrorResponse(c, &errResp)
		return
	}

	response.SendSuccessResponse(c, http.StatusCreated, nil, nil)

}

func (usr *userHandler) Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		usr.logger.Error("failed to bind login request body", zap.Error(err))
		response.SendErrorResponse(c, &response.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	token, errResp := usr.userService.Login(req)
	if errResp.Message != "" {
		response.SendErrorResponse(c, &errResp)
		return
	}

	response.SendSuccessResponse(c, http.StatusOK, token, nil)

}
