package service

import "user-auth-service/models"

type User interface {
	Register(req models.RegisterUser) error
}
