package routing

import (
	"portfolio-go/handler"
	"portfolio-go/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeRoute() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"*"},
	}))

	r.Use(cors.Default())
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

	r.GET("/userprofile", handler.GetUserProfile)

	socialMedia := r.Group("/socialmedia")
	{
		socialMedia.Use(middleware.AuthMiddleware())
		socialMedia.GET("/:id", handler.GetDetailSocialMedia)
		socialMedia.POST("/", handler.CreateSocialMedia)
		socialMedia.PUT("/", handler.UpdateSocialMedia)
		socialMedia.DELETE("/:id", handler.DeleteSocialMedia)
		socialMedia.GET("/user", handler.GetSocialMediasByUser)
	}

	r.SetTrustedProxies(nil)

	return r
}
