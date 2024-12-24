package middleware

import (
	"log"
	"net/http"

	"github.com/brotigen23/gopherMart/internal/utils"
)

var JWTSecretKey = "secret"

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			log.Println(err.Error())
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		}
		// * Если ошибка то токен непонятный, иначе все норм
		user, err := utils.GetUserLoginFromJWT(cookie.Value, JWTSecretKey)
		if err != nil {
			log.Println(err.Error())
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		}
		r.AddCookie(&http.Cookie{Name: "userLogin", Value: user})
	})
}

func Auth(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("token")
			if err != nil {
				log.Println(err.Error())
				http.Error(rw, err.Error(), http.StatusUnauthorized)
				return
			}
			// * Если ошибка то токен непонятный, иначе все норм
			user, err := utils.GetUserLoginFromJWT(cookie.Value, secretKey)
			if err != nil {
				log.Println(err.Error())
				http.Error(rw, err.Error(), http.StatusUnauthorized)
				return
			}
			r.AddCookie(&http.Cookie{Name: "userLogin", Value: user})
			next.ServeHTTP(rw, r)
		})
	}
}
