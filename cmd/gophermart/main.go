package main

import (
	"log"

	"github.com/brotigen23/gopherMart/internal/app"
)

func main() {
	err := app.Run(":8080")
	if err != nil {
		log.Println(err)
		return
	}
}
