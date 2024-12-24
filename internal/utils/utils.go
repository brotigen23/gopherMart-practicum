package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/brotigen23/gopherMart/internal/dto"
	"github.com/golang-jwt/jwt/v4"
)

type UserJWTClaims struct {
	jwt.RegisteredClaims
	Login string
}

func UnmarhallUserJWT(r io.ReadCloser) (*dto.User, error) {
	var user dto.User
	var buffer bytes.Buffer
	_, err := buffer.ReadFrom(r)
	if err != nil {
		// TODO: обработать ошибку
		log.Printf("error: %v", err.Error())
		return nil, err
	}
	if err = json.Unmarshal(buffer.Bytes(), &user); err != nil {
		log.Printf("error: %v", err.Error())
		return nil, err
	}
	return &user, nil
}

func BuildJWTString(login string, key string, expires time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserJWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expires)),
		},
		Login: login,
	})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func GetUserLoginFromJWT(tokenString string, key string) (string, error) {
	claims := &UserJWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("token is invalid")
	}

	return claims.Login, nil
}

func IsOrderCorrect(order string) bool {
	o, err := strconv.Atoi(order)
	if err != nil {
		return false
	}
	log.Println(o)
	return true
}
