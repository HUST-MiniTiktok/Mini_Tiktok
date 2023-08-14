package crypt

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt the password by bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(hashedPassword), err
}

// VerifyPassword Verify the password is consistent with the hashed password in the database
func VerifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}