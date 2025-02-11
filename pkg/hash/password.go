package hash

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// GENERATE PASSWORD
func HashingPassword(password string) (string, error) {

	passwordBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	if err != nil {
		return "", errors.New("failed to generate hashed password")
	}
	return string(hashedPassword), err
}

// VALIDATE PASSWORD
func ValidatePassword(incomingPassword, encriptedUserPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encriptedUserPassword), []byte(incomingPassword))
	return err == nil
}
