package main

import (
	"log"
	"os"
	"portfolio-go/routing"

	"github.com/gin-gonic/gin"
)

func main() {

	if os.Getenv("ENV") != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := routing.InitializeRoute()
	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Print(err)
	}
	r.Run()
}
