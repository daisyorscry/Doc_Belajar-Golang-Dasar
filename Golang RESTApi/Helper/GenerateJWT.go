package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("1234")

func GenerateJWT(userId int, username string) (string, error) {

	claims := &jwt.MapClaims{
		"userId":   userId,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
