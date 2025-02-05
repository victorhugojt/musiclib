package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const apiKey = "supersecret"

func GenerateToken(email, user_id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(apiKey))
}
