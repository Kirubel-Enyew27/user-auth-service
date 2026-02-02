package user

import (
	"net/http"
	"user-auth-service/handler"
	"user-auth-service/models"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := usr.userService.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "user created successfully"})
}
