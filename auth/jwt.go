package auth

import (
	"api/config"
	"api/utils/console"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CreateJWTToken(userId uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SECRET_KEY)
}

func ValidateToken(r *http.Request) error {
	tokenString := ExtractJWTToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.SECRET_KEY, nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		console.Pretty(claims)
	}
	return nil
}

func ExtractTokenID(r *http.Request) (uint32, error) {
	tokenString := ExtractJWTToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.SECRET_KEY, nil
	})

	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if !ok {
			return 0, nil
		}
		uid, err := strconv.ParseUint(
			fmt.Sprintf("%.0f", claims["userId"]), 10, 32)
		if err != nil {
			return 0, nil
		}
		return uint32(uid), nil
	}
	return 0, nil
}

func ExtractJWTToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	fmt.Println("bearerToken: ", bearerToken)
	return ""
}