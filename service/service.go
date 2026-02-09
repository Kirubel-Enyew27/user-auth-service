package service

import (
	"user-auth-service/models"
	"user-auth-service/pkg/response"
)

type User interface {
	Register(req models.RegisterUser) response.ErrorResponse
	Login(req models.LoginRequest) (string, response.ErrorResponse)
	ChangePassword(req models.ChangePassword) response.ErrorResponse
	SuspendUser(id string) response.ErrorResponse
	ActivateUser(id string) response.ErrorResponse
}
