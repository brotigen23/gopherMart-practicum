package app

import (
	"log"

	"github.com/brotigen23/gopherMart/internal/config"
	"github.com/brotigen23/gopherMart/internal/server"
	"github.com/joho/godotenv"
)

func Run(serverAddr string) error {
	godotenv.Load()
	config, err := config.NewConfig()
	if err != nil {
		return err
	}
	log.Println(config)
	server := server.NewServer(config)
	err = server.Run()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
