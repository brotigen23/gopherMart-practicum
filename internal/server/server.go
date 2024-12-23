package server

import (
	"log"
	"net/http"

	"github.com/brotigen23/gopherMart/internal/config"
	"github.com/brotigen23/gopherMart/internal/handler"
	"github.com/brotigen23/gopherMart/internal/middleware"
	"github.com/brotigen23/gopherMart/internal/repository"
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
	userRepository, err := repository.NewPostgresUserRepository("postgres", "host=localhost port=5432 user=myuser password=1234 dbname=mydb sslmode=disable")
	if err != nil {
		return err
	}
	userRepository.GetUserByLogin()

	// TODO: создание сервисов
	log.Println("Services")
	// TODO: создание хендлеров
	log.Println("Handlers")
	userHandler := handler.NewUserHandler(nil)

	// TODO: создание роутера
	log.Println("Router")
	router := chi.NewRouter()
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.Compress(3, "plain/text", "application/json"))

	router.With(middleware.ValidateUser).Post("/api/user/register", userHandler.Register)
	router.With(middleware.ValidateUser).Post("/api/user/login", userHandler.Login)

	router.With(middleware.Auth).Post("/api/user/orders", userHandler.SaveOrder)
	router.With(middleware.Auth).Get("/api/user/orders", userHandler.GetOrders)

	router.With(middleware.Auth).Get("/api/user/balance", userHandler.GetBalance)
	router.With(middleware.Auth).Post("/api/user/balance/withdraw", userHandler.Withdraw)

	router.With(middleware.Auth).Get("/api/user/withdrawals", userHandler.GetWithDrawals)

	serv := &http.Server{
		Addr:    s.config.Server.Host + ":" + s.config.Server.Port,
		Handler: router,
	}

	return serv.ListenAndServe()
}
