package app

import (
	"log"

	"github.com/brotigen23/gopherMart/internal/config"
	"github.com/brotigen23/gopherMart/internal/server"
)

func Run(serverAddr string) error {
	config, err := config.NewConfig("config.yaml")
	if err != nil {
		return err
	}

	server := server.NewServer(config)
	err = server.Run()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
