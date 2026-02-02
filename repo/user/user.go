package user

import (
	"database/sql"
	"fmt"
	"user-auth-service/models"
	"user-auth-service/repo"

	"go.uber.org/zap"
)

type userRepo struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewRepo(db *sql.DB, logger *zap.Logger) repo.User {
	return &userRepo{
		db:     db,
		logger: logger,
	}
}
func (usr *userRepo) Register(user models.User) error {
	createUserSql := `
	INSERT INTO users(id, username, password, phone, email, role, status, created_at)
	VALUES (&1, &2, &3, &4, &5, &6, &7, &8)
	`

	_, err := usr.db.Exec(createUserSql, user.ID, user.Username, user.Password, user.Phone, user.Email, user.Role, user.Status, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil

}

func (usr *userRepo) GetUserByNameOrPhone(username, phone string) (models.User, error) {
	row, err := usr.db.Query("SELECT * FROM users WHERE username=? OR phone=?", username, phone)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get user by username")
	}
	defer row.Close()

	var user models.User

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Phone, &user.Email, &user.Role, &user.Status, &user.CreatedAt)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}


