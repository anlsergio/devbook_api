package security

import "golang.org/x/crypto/bcrypt"

// Hash - receives a string and turns it into a hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CheckPassword - Checks if a provided plain text password and a hash match
func CheckPassword(hashPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))
}