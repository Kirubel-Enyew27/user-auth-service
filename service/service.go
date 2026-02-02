package service

import (
	"user-auth-service/models"
	"user-auth-service/pkg/response"
)

type User interface {
	Register(req models.RegisterUser) response.ErrorResponse
}
