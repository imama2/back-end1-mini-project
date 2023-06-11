package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	saltCost = 12
)

// HashPass menghasilkan hash dari password yang diberikan.
func HashPass(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), saltCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

// ComparePass membandingkan password biasa dengan hash password yang ada.
// Mengembalikan nilai true jika password cocok dengan hash,
// atau false jika tidak cocok atau terjadi kesalahan.
func ComparePass(hashPass, password []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	return err == nil
}
