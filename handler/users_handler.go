package handler

import (
	"log"
	"net/http"
	"portfolio-go/request"
	"portfolio-go/response"
	"portfolio-go/service"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := service.GetUsers(c)
	if err != nil {
		response.ErrorResponse(
			c,
			http.StatusUnprocessableEntity,
			err.Error(),
		)
		log.Fatalln("error get data users: ", err.Error())
		return
	}

	response.SuccessResponse(c, response.BaseResponse{
		Data:       users,
		Message:    "successfully get users",
		StatusCode: 200,
	})
}

func Register(c *gin.Context) {
	var request request.UserRequest

	if err := c.ShouldBind(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
		return
	}

	user, err := service.Register(c, request)
	if err != nil {
		response.ErrorResponse(
			c,
			http.StatusUnprocessableEntity,
			err.Error(),
		)
		log.Fatalln("error register: ", err.Error())
		return
	}

	response.SuccessResponse(c, response.BaseResponse{
		Data:       user,
		Message:    "successfully registered",
		StatusCode: 200,
	})
}
