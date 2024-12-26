package main

import (
	"log"

	"github.com/brotigen23/gopherMart/internal/utils"
)

func main() {
	//app.Run(":8080")
	log.Println(utils.IsOrderCorrect("38215667007"))
	// 3 *8 2 *1 5 *6 6 *7 0 *0 7
	// 3 *7 2 *2 5 *3 6 *5 0 *0 7
}
