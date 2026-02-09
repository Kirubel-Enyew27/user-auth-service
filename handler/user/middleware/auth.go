package middleware

import (
	"net/http"
	"strings"
	"user-auth-service/pkg/response"
	"user-auth-service/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		response.SendErrorResponse(c, &response.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "token is not provided",
		})
		return
	}

	token := strings.Split(authHeader, "Bearer")[1]
	err := utils.ValidateToken(token)
	if err != nil {
		response.SendErrorResponse(c, &response.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
		return
	}

	c.Next()
}
