package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	const cost = 14
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
