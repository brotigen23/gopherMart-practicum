package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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
	expires := time.Hour * 1024
	// TODO: вынести секретный ключ в переменную окружения
	JWTSecretKey := "secret"
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
	user, err := r.Cookie("userLogin")
	if err != nil {
		log.Printf("error: %v", ErrJWT.Error())
		http.Error(rw, ErrJWT.Error(), http.StatusUnauthorized)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// TODO: обработать ошибку
		log.Printf("error: %v", err.Error())
		http.Error(rw, ErrBadRequest.Error(), http.StatusBadRequest)
		return
	}
	order := string(body)
	if !utils.IsOrderCorrect(order) {
		log.Printf("error: %v", ErrJWT.Error())
		http.Error(rw, ErrJWT.Error(), http.StatusBadRequest)
		return
	}
	h.userService.SaveOrder(user.Value, order)
	rw.WriteHeader(http.StatusAccepted)
	// create goroutine to check order status
}

func (h *userHandler) GetOrders(rw http.ResponseWriter, r *http.Request) {
	// Get user's orders
	userLogin, err := r.Cookie("userLogin")
	if err != nil {
		log.Printf("error: %v", ErrJWT.Error())
		http.Error(rw, ErrJWT.Error(), http.StatusBadRequest)
		return
	}
	orders, err := h.userService.GetOrders(userLogin.Value)
	if err != nil {
		log.Printf("error: %v", ErrBadRequest.Error())
		http.Error(rw, ErrBadRequest.Error(), http.StatusNoContent)
		return
	}
	for i, order := range orders {
		resp, err := http.Get(h.Config.AccrualSystemAddress + "/api/orders/" + order.Number)
		// Check order's status
		if err != nil {
			log.Printf("error: %v", ErrAccrualSystem.Error())
			http.Error(rw, ErrAccrualSystem.Error(), http.StatusBadRequest)
			return
		}
		o, err := utils.UnmarhallOrder(resp.Body)
		if err != nil {
			log.Printf("error: %v", ErrAccrualSystem.Error())
			http.Error(rw, ErrAccrualSystem.Error(), http.StatusBadRequest)
			return
		}
		orders[i].Status = o.Status
		orders[i].Accrual = o.Accrual
	}
	resp, err := json.Marshal(orders)
	if err != nil {
		log.Printf("error: %v", ErrAccrualSystem.Error())
		http.Error(rw, ErrAccrualSystem.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(resp)
}

func (h *userHandler) GetBalance(rw http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("userLogin")
	if err != nil {
		return
	}

}

func (h *userHandler) Withdraw(rw http.ResponseWriter, r *http.Request) {

}

func (h *userHandler) GetWithDrawals(rw http.ResponseWriter, r *http.Request) {

}
