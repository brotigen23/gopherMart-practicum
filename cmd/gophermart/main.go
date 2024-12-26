package main

import (
	"log"

	"github.com/brotigen23/gopherMart/internal/app"
	"github.com/brotigen23/gopherMart/internal/utils"
)

//55501241
// 1 0 4 2
func main() {
	log.Println(utils.IsOrderCorrect("5675110315"))
	app.Run(":8080")
}
