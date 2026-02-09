package repo

import (
	"user-auth-service/models"
	"user-auth-service/pkg/response"
)

type User interface {
	Register(user models.User) response.ErrorResponse
	GetUserByNameOrPhone(username, phone string) (models.User, response.ErrorResponse)
	GetUserByID(id string) (models.User, response.ErrorResponse)
	UpdatePassword(id, newPasswordHash string) response.ErrorResponse
	SuspendUser(id string) response.ErrorResponse
	ActivateUser(id string) response.ErrorResponse
}
