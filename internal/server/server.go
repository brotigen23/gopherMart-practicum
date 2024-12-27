package server

import (
	"log"
	"net/http"

	"github.com/brotigen23/gopherMart/internal/config"
	"github.com/brotigen23/gopherMart/internal/handler"
	"github.com/brotigen23/gopherMart/internal/middleware"
	"github.com/brotigen23/gopherMart/internal/repository"
	"github.com/brotigen23/gopherMart/internal/service"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s Server) Run() error {
	// TODO: создание подключения к БД
	log.Println("Repository")
	userRepository, err := repository.NewPostgresUserRepository("postgres", s.config.DatabaseURI)
	if err != nil {
		return err
	}

	// TODO: создание сервисов
	log.Println("Services")
	userService := service.NewUserService(userRepository)

	// TODO: создание хендлеров
	log.Println("Handlers")
	userHandler := handler.NewUserHandler(userService, s.config)

	// TODO: создание роутера
	log.Println("Router")
	router := chi.NewRouter()
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.Compress(3, "plain/text", "application/json"))

	router.Post("/api/user/register", userHandler.Register)
	router.Post("/api/user/login", userHandler.Login)

	router.With(middleware.Auth(s.config.JWTSecretKey)).Post("/api/user/orders", userHandler.SaveOrder)
	router.With(middleware.Auth(s.config.JWTSecretKey)).Get("/api/user/orders", userHandler.GetOrders)

	router.With(middleware.Auth(s.config.JWTSecretKey)).Get("/api/user/balance", userHandler.GetBalance)
	router.With(middleware.Auth(s.config.JWTSecretKey)).Post("/api/user/balance/withdraw", userHandler.Withdraw)

	router.With(middleware.Auth(s.config.JWTSecretKey)).Get("/api/user/withdrawals", userHandler.GetWithdrawals)

	serv := &http.Server{
		Addr:    s.config.RunAdress,
		Handler: router,
	}

	return serv.ListenAndServe()
}
