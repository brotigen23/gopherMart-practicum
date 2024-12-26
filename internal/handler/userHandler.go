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
	user, err := utils.UnmarhallUser(r.Body)
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
	jwtString, err := utils.BuildJWTString(user.Login, h.Config.JWTSecretKey, expires)
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
	http.SetCookie(rw, cookie)
	rw.WriteHeader(http.StatusOK)
}

func (h *userHandler) Login(rw http.ResponseWriter, r *http.Request) {
	user, err := utils.UnmarhallUser(r.Body)
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
	expires := time.Hour * 1024
	jwtString, err := utils.BuildJWTString(user.Login, h.Config.JWTSecretKey, expires)
	if err != nil {
		log.Printf("error: %v", err.Error())
		return
	}

	// Отправляем на базу
	cookie := &http.Cookie{
		Name:  "token",
		Value: jwtString,
	}
	http.SetCookie(rw, cookie)
	rw.WriteHeader(http.StatusOK)
}

func (h *userHandler) SaveOrder(rw http.ResponseWriter, r *http.Request) {
	userLogin, err := r.Cookie("userLogin")
	if err != nil {
		log.Printf("error: %v", err)
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
		log.Printf("error: %v", ErrBadOrderNumber.Error())
		http.Error(rw, ErrBadOrderNumber.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = h.userService.SaveOrder(userLogin.Value, order)
	switch err {
	case nil:
		rw.WriteHeader(http.StatusAccepted)
	case service.ErrOrderAlreadySave:
		http.Error(rw, service.ErrOrderAlreadySave.Error(), http.StatusOK)
		return
	case service.ErrOrderSavedByOtherUser:
		http.Error(rw, service.ErrOrderSavedByOtherUser.Error(), http.StatusConflict)
		return
	default:
		log.Println(err.Error())
		return
	}
	log.Println("order", order, "registered by", userLogin)
	// create goroutine to check order status
	for {
		time.Sleep(time.Second)
		resp, err := http.Get(h.Config.AccrualSystemAddress + "/api/orders/" + order)
		if err != nil {
			log.Println(err.Error())
			return
		}
		if resp.StatusCode != http.StatusOK {
			continue
		}
		o, err := utils.UnmarhallOrder(resp.Body)
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer resp.Body.Close()
		log.Println(o)
		switch o.Status {
		case "PROCESSED":
			log.Println("PROCESSED")
			h.userService.UpdateUserBalance(userLogin.Value, o.Accrual)
			return
		default:
			return
		}
	}
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
		if resp.StatusCode != http.StatusOK {
			log.Println("order", order.Number, "is not registered")
			continue
		}
		log.Println("status code:", resp.StatusCode)
		// Check order's status
		if err != nil {
			http.Error(rw, ErrAccrualSystem.Error(), http.StatusBadRequest)
			return
		}
		o, err := utils.UnmarhallOrder(resp.Body)
		if err != nil {
			http.Error(rw, ErrAccrualSystem.Error(), http.StatusBadRequest)
			return
		}
		orders[i].Status = o.Status
		orders[i].Accrual = o.Accrual
	}
	resp, err := json.Marshal(orders)
	if err != nil {
		http.Error(rw, ErrAccrualSystem.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(resp)
}

func (h *userHandler) GetBalance(rw http.ResponseWriter, r *http.Request) {
	userLogin, err := r.Cookie("userLogin")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	balance, err := h.userService.GetUserBalance(userLogin.Value)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := json.Marshal(balance)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(resp)

}

func (h *userHandler) Withdraw(rw http.ResponseWriter, r *http.Request) {

}

func (h *userHandler) GetWithDrawals(rw http.ResponseWriter, r *http.Request) {

}
