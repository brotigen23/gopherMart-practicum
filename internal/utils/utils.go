package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/brotigen23/gopherMart/internal/dto"
	"github.com/golang-jwt/jwt/v4"
)

type UserJWTClaims struct {
	jwt.RegisteredClaims
	Login string
}

func UnmarhallUser(r io.ReadCloser) (*dto.User, error) {
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
func UnmarhallOrder(r io.ReadCloser) (*dto.Order, error) {
	var order dto.Order
	var buffer bytes.Buffer
	_, err := buffer.ReadFrom(r)
	if err != nil {
		// TODO: обработать ошибку
		log.Printf("error: %v", err.Error())
		return nil, err
	}
	if err = json.Unmarshal(buffer.Bytes(), &order); err != nil {
		log.Printf("error: %v", err.Error())
		return nil, err
	}
	return &order, nil
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

// 38215667007
func IsOrderCorrect(order string) bool {
	n := len(order)
	number := 0
	result := 0
	for i := 0; i < n; i++ {
		number = int(order[i]) - '0'
		if (i+1)%2 != 0 {
			result += number
			continue
		}
		number *= 2
		if number > 9 {
			number = number/10 + number%10
		}
		result += number
	}
	return result%10 == 0
}
