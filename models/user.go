package models

import validation "github.com/go-ozzo/ozzo-validation"

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
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

func (usr *User) Validate() error {
	err := validation.ValidateStruct(usr,
		validation.Field(usr.Username, validation.Required, validation.Length(3, 0)),
		validation.Field(usr.Password, validation.Required, validation.Length(4, 8)),
		validation.Field(usr.Phone, validation.Required),
		validation.Field(usr.Email, validation.Required),
	)

	return err
}
