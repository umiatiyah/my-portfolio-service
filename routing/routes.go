package routing

import (
	"portfolio-go/handler"
	"portfolio-go/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoute() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handler.CheckHealth)

	auth := r.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/register", handler.Register)
	}

	users := r.Group("/users")
	{
		users.Use(middleware.AuthMiddleware())
		users.GET("/", handler.GetUsers)
	}

	r.SetTrustedProxies(nil)

	return r
}
