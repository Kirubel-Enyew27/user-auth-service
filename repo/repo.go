package repo

import (
	"user-auth-service/models"
	"user-auth-service/pkg/response"
)

type User interface {
	Register(user models.User) response.ErrorResponse
	GetUserByNameOrPhone(username, phone string) (models.User, response.ErrorResponse)
}
