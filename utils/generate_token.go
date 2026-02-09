package utils

import (
	"time"
	"user-auth-service/models"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user-data": user,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(viper.GetString("JWT_KEY"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
