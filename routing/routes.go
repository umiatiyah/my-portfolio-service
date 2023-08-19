package routing

import (
	"portfolio-go/handler"
	"portfolio-go/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoute() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handler.CheckHealth)

	users := r.Group("/users")
	{
		users.Use(middleware.AuthMiddleware())
		users.GET("/", handler.GetUsers)
	}

	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)

	return r
}
