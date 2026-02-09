package handler

import "github.com/gin-gonic/gin"

type User interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}
