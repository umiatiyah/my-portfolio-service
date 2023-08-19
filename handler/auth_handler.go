package handler

import (
	"log"
	"net/http"
	"portfolio-go/request"
	"portfolio-go/response"
	"portfolio-go/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request request.AuthRequest

	if err := c.ShouldBind(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
		return
	}

	token, err := service.Login(c, request)
	if err != nil {
		response.ErrorResponse(
			c,
			http.StatusUnprocessableEntity,
			err.Error(),
		)
		log.Println("error login: ", err.Error())
		return
	}

	response.SuccessResponse(c, response.BaseResponse{
		Data:       token,
		Message:    "successfully login",
		StatusCode: 200,
	})
}
