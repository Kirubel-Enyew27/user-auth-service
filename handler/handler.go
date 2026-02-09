package handler

import "github.com/gin-gonic/gin"

type User interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	ChangePassword(c *gin.Context)
	SuspendUser(c *gin.Context)
	ActivateUser(c *gin.Context)
}
