package utils

import (
	"fmt"
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

func ValidateToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return viper.GetString("JWt_KEY"), nil
	})
	if err != nil {
		return fmt.Errorf("error parsing token: %w", err)
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
