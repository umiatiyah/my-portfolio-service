package handler

import (
	"log"
	"net/http"
	"portfolio-go/response"
	"portfolio-go/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	userid := c.Param("userId")
	id, err := strconv.Atoi(userid)
	users, err := service.GetUserProfile(c, id)
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
