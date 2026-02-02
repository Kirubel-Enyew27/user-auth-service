package repo

import "user-auth-service/models"

type User interface {
	Register(user models.User) error
}
