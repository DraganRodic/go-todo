package utils

import (
	"time"
	"todo-api/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func getSecret() []byte {
	return []byte(config.AppConfig.JWTSecret)
}

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getSecret())
}

func ParseToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return getSecret(), nil
	})
}
