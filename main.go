package main

import (
	"log"
	"os"
	"portfolio-go/routing"
)

func main() {

	r := routing.InitializeRoute()
	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Print(err)
	}
	r.Run()
}
