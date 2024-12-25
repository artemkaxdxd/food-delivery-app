package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateJwtToken(claims jwt.Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GenerateAPIToken(id uint, secret string, expiration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"iat": time.Now().Unix() - 10,
		"exp": time.Now().Add(expiration).Unix(),
	}
	return generateJwtToken(claims, secret)
}
