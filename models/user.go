package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Role string
type Status string

const (
	RoleUser       Role = "user"
	RoleAdmin      Role = "admin"
	RoleSuperAdmin Role = "super-admin"

	StatusActive   Status = "active"
	StatusPending  Status = "pending"
	StatusInActive Status = "inactive"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (usr *RegisterUser) Validate() error {
	err := validation.ValidateStruct(usr,
		validation.Field(usr.Username, validation.Required, validation.Length(3, 50)),
		validation.Field(usr.Password, validation.Required, validation.Length(4, 8)),
		validation.Field(usr.Phone, validation.Required),
		validation.Field(usr.Email, validation.Required, is.Email),
		validation.Field(usr.Role, validation.Required),
	)

	return err
}

func (auth *LoginRequest) Validate() error {
	return validation.ValidateStruct(auth,
		validation.Field(auth.Username, validation.Required),
		validation.Field(auth.Password, validation.Required),
	)
}
