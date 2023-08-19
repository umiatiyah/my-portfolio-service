package handler

import (
	"log"
	"portfolio-go/service"

	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
	err := service.CheckHealth(c)
	if (err != nil){
		log.Fatalln("error checking health")
	}
	c.JSON(200, gin.H{
		"message": err,
	})
}