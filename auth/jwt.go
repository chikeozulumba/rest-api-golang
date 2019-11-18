package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

func CreateJWTToken(userId uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

	})
}