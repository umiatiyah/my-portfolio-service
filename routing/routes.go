package routing

import (
	"portfolio-go/handler"
	"portfolio-go/middleware"

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
		users.Use(middleware.AuthMiddleware())
		// Define routes within the group
		users.GET("/", handler.GetUsers)
	}

	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)

	r.Run() // listen and serve on 0.0.0.0:8080
}
