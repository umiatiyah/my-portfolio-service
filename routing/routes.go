package routing

import (
	"portfolio-go/handler"

	"github.com/gin-gonic/gin"
)

func InitializeRoute() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/health", handler.CheckHealth)

	users := r.Group("/users")
	{
		// Define routes within the group
		users.GET("/", handler.GetUsers)
		users.POST("/login", handler.Login)
		users.POST("/register", handler.Register)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}
