package handler

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/brotigen23/gopherMart/internal/config"
	"github.com/brotigen23/gopherMart/internal/service"
	"github.com/brotigen23/gopherMart/internal/utils"
)

type userHandler struct {
	Config *config.Config

	userService *service.UserService
}

func NewUserHandler(userService *service.UserService, config *config.Config) *userHandler {
	return &userHandler{
		Config:      config,
		userService: userService,
	}
}

func (h *userHandler) Register(rw http.ResponseWriter, r *http.Request) {
	log.Println("Register handler")
	user, err := utils.UnmarhallUserJWT(r.Body)
	if err != nil {
		log.Printf("error: %v", err.Error())
		http.Error(rw, ErrInternalServer.Error(), http.StatusInternalServerError)
		return
	}

	// Проверяем наличие логина в БД
	if h.userService.IsUserExists(user.Login) {
		log.Printf("error: %v", ErrUserExists)
		http.Error(rw, ErrUserExists.Error(), http.StatusConflict)
		return
	}

	// Регистрируем новую запись в БД
	err = h.userService.SaveUser(user.Login, user.Password)
	if err != nil {
		log.Printf("error: %v", err.Error())
		http.Error(rw, ErrInternalServer.Error(), http.StatusInternalServerError)
		return
	}

	// Создаем JWT токен
	expires := time.Minute * 15
	// TODO: вынести секретный ключ в переменную окружения
	JWTSecretKey := "secret_key"
	jwtString, err := utils.BuildJWTString(user.Login, JWTSecretKey, expires)
	if err != nil {
		log.Printf("error: %v", err.Error())
		http.Error(rw, ErrInternalServer.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем на базу
	cookie := &http.Cookie{
		Name:  "token",
		Value: jwtString,
	}
	r.AddCookie(cookie)
	http.SetCookie(rw, cookie)
	rw.WriteHeader(http.StatusOK)
}

func (h *userHandler) Login(rw http.ResponseWriter, r *http.Request) {
	user, err := utils.UnmarhallUserJWT(r.Body)
	if err != nil {
		http.Error(rw, "server error", http.StatusInternalServerError)
		return
	}
	//* Проверяем наличие логина в БД и правильность введенного пароля
	if !h.userService.IsUserExists(user.Login) {
		log.Printf("error: %v", ErrWrongLogin.Error())
		http.Error(rw, ErrWrongLogin.Error(), http.StatusUnauthorized)
		return
	}
	password, err := h.userService.GetUserPasswordByLogin(user.Login)
	if err != nil {
		log.Printf("error: %v", err.Error())
		http.Error(rw, err.Error(), http.StatusUnauthorized)
		return
	}
	if user.Password != password {
		log.Printf("error: %v", ErrWrongPassword.Error())
		http.Error(rw, ErrWrongPassword.Error(), http.StatusUnauthorized)
		return
	}
	// Создаем JWT токен
	expires := time.Minute * 15
	// TODO: вынести секретный ключ в переменную окружения
	JWTSecretKey := "secret_key"
	jwtString, err := utils.BuildJWTString(user.Login, JWTSecretKey, expires)
	if err != nil {
		log.Printf("error: %v", err.Error())
		return
	}

	// Отправляем на базу
	cookie := &http.Cookie{
		Name:  "token",
		Value: jwtString,
	}
	r.AddCookie(cookie)
	http.SetCookie(rw, cookie)
	rw.WriteHeader(http.StatusOK)
}

func (h *userHandler) SaveOrder(rw http.ResponseWriter, r *http.Request) {
	user, err := utils.UnmarhallUserJWT(r.Body)
	if err != nil {
		log.Printf("error: %v", ErrJWT.Error())
		http.Error(rw, ErrJWT.Error(), http.StatusUnauthorized)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// TODO: обработать ошибку
		log.Printf("error: %v", err.Error())
		return
	}
	order, err := strconv.Atoi(string(body))
	if err != nil {
		// TODO: обработать ошибку
		log.Printf("error: %v", err.Error())
		return
	}
	if !utils.IsOrderCorrect(order) {
		log.Printf("error: %v", ErrJWT.Error())
		http.Error(rw, ErrJWT.Error(), http.StatusUnauthorized)
		return
	}
	h.userService.SaveOrder(user.Login, order)
}

func (h *userHandler) GetOrders(rw http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(h.Config.AccrualSystemAddress)
	if err != nil {
		return
	}
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Println(string(buf))
}

func (h *userHandler) GetBalance(rw http.ResponseWriter, r *http.Request) {
	userLogin, err := r.Cookie("userLogin")
	if err != nil {
		return
	}
	log.Println(userLogin)
}

func (h *userHandler) Withdraw(rw http.ResponseWriter, r *http.Request) {

}

func (h *userHandler) GetWithDrawals(rw http.ResponseWriter, r *http.Request) {

}
