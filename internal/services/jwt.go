package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT secret key
var jwtSecret = os.Getenv("JWT_SECRET")

// GenerateJWT creates a JWT token for authentication
func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // Expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// ValidateJWT checks if a token is valid
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
