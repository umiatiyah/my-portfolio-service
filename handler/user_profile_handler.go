package handler

import (
	"log"
	"net/http"
	"portfolio-go/response"
	"portfolio-go/service"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	}
	users, err := service.GetUserProfile(c, username)
	if err != nil {
		response.ErrorResponse(
			c,
			http.StatusUnprocessableEntity,
			err.Error(),
		)
		log.Println("error get data users: ", err.Error())
		return
	}

	response.SuccessResponse(c, response.BaseResponse{
		Data:       users,
		Message:    "successfully get users",
		StatusCode: 200,
	})
}
