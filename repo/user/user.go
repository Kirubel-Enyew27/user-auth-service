package user

import (
	"database/sql"
	"net/http"
	"user-auth-service/models"
	"user-auth-service/pkg/response"
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
func (usr *userRepo) Register(user models.User) response.ErrorResponse {
	createUserSql := `
	INSERT INTO users(id, username, password, phone, email, role, status, created_at)
	VALUES (&1, &2, &3, &4, &5, &6, &7, &8)
	`

	_, err := usr.db.Exec(createUserSql, user.ID, user.Username, user.Password, user.Phone, user.Email, user.Role, user.Status, user.CreatedAt)
	if err != nil {
		usr.logger.Error("failed to create user account", zap.Error(err))
		return response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to create user account",
		}
	}

	return response.ErrorResponse{}

}

func (usr *userRepo) GetUserByNameOrPhone(username, phone string) (models.User, response.ErrorResponse) {
	row, err := usr.db.Query("SELECT * FROM users WHERE username=? OR phone=?", username, phone)
	if err != nil {
		usr.logger.Error("failed to get user by phone or username", zap.Error(err))
		return models.User{}, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to get user",
		}
	}
	defer row.Close()

	var user models.User

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Phone, &user.Email, &user.Role, &user.Status, &user.CreatedAt)
	if err != nil {
		usr.logger.Error("failed to scan user rows", zap.Error(err))
		return models.User{}, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to scan user rows",
		}
	}

	return user, response.ErrorResponse{}
}

func (usr *userRepo) GetUserByID(id string) (models.User, response.ErrorResponse) {
	row, err := usr.db.Query("SELECT * FROM users WHERE id=?", id)
	if err != nil {
		usr.logger.Error("failed to get user by id", zap.Error(err))
		return models.User{}, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to get user",
		}
	}
	defer row.Close()

	var user models.User

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Phone, &user.Email, &user.Role, &user.Status, &user.CreatedAt)
	if err != nil {
		usr.logger.Error("failed to scan user rows", zap.Error(err))
		return models.User{}, response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to scan user rows",
		}
	}

	return user, response.ErrorResponse{}
}

func (usr *userRepo) UpdatePassword(id, newPasswordHash string) response.ErrorResponse {
	_, err := usr.db.Exec("UPDATE users SET password=? WHERE id=?", newPasswordHash, id)
	if err != nil {
		usr.logger.Error("failed to update password", zap.Error(err))
		return response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	return response.ErrorResponse{}
}

func (usr *userRepo) SuspendUser(id string) response.ErrorResponse {
	_, err := usr.db.Exec("UPDATE users SET status=? WHERE id=?", models.StatusPending, id)
	if err != nil {
		usr.logger.Error("failed to suspend user", zap.Error(err))
		return response.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	return response.ErrorResponse{}
}
