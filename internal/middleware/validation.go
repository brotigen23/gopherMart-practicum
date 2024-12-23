package middleware

import (
	"log"
	"net/http"

	"github.com/brotigen23/gopherMart/internal/utils"
)

func ValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		body := r.Clone(r.Context())

		if body.Header.Get("Content-Type") != "application/json" {
			log.Printf("error: %v", "not json")
			http.Error(rw, ErrContentType.Error(), http.StatusBadRequest)
			return
		}

		user, err := utils.UnmarhallUser(body.Body)
		if err != nil {
			log.Printf("error: %v", err.Error())
			http.Error(rw, ErrNotValidJSON.Error(), http.StatusBadRequest)
			return
		}
		if user.Login == "" || user.Password == "" {
			log.Printf("error: %v", "empty values")
			http.Error(rw, ErrNotValidJSON.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(rw, r)
	})
}

func ValitadeContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		switch r.Header.Get("Content-Type") {
		case "plaint/text":
			switch r.URL.Path {
			case "/api/user/login":
				return
			default:
				return
			}
		case "application/json":
			switch r.URL.Path {

			}
		}
		switch r.URL.Path {

		case "/api/user/login", "/api/user/register":
			switch r.Header.Get("Content-Type") {
			case "application/json":
				break
			}

		case "/api/user/orders":
			break
		}
		next.ServeHTTP(rw, r)
	})

}
