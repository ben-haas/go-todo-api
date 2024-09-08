package util

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes user password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash validates user password hash
func CheckPasswordHash(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}
