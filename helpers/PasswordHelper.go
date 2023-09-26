package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHelper struct{}

func (ph *PasswordHelper) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	return string(bytes), err
}

func (ph *PasswordHelper) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
