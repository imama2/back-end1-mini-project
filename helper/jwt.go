package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

const secretKey = "access-login"

// GenerateToken menghasilkan token JWT dengan mengambil id pengguna dan username.
func GenerateToken(id uint, username string) (string, error) {
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken memverifikasi keaslian token JWT yang diberikan di header Authorization.
func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in to process")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errResponse
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errResponse
	}

	return claims, nil
}
