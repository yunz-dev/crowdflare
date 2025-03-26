package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// User structure
type User struct {
	Username string
	Password string
}

// Mock user storage (replace with DB later)
var users = make(map[string]string)

// HashPassword securely hashes a password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword verifies a hashed password
func CheckPassword(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}

// RegisterUser saves a user in the mock DB
func RegisterUser(username, password string) error {
	if _, exists := users[username]; exists {
		return errors.New("user already exists")
	}
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}
	users[username] = hashedPassword
	return nil
}

// AuthenticateUser checks user credentials
func AuthenticateUser(username, password string) (bool, error) {
	hashedPassword, exists := users[username]
	if !exists {
		return false, errors.New("user not found")
	}
	err := CheckPassword(hashedPassword, password)
	return err == nil, err
}
