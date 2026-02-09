package router

import (
	"user-auth-service/handler"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine, userHandler handler.User) {
	router.POST("/user/register", userHandler.Register)
	router.POST("/user/login", userHandler.Login)
}
