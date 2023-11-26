package handler

import (
	"fmt"
	"log"
	"net/http"
	"portfolio-go/model"
	"portfolio-go/response"
	"portfolio-go/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(c *gin.Context) {
	var request model.SocialMedia

	if err := c.ShouldBind(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
		return
	}

	socialMedia, err := service.CreateSocialMedia(c, request)
	if err != nil {
		response.ErrorResponse(
			c,
			http.StatusUnprocessableEntity,
			err.Error(),
		)
		log.Println("error create social media: ", err.Error())
		return
	}

	response.SuccessResponse(c, response.BaseResponse{
		Data:       socialMedia,
		Message:    "successfully created social media",
		StatusCode: 200,
	})
}

func UpdateSocialMedia(c *gin.Context) {
	var request model.SocialMedia

	if err := c.ShouldBind(&request); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
		return
	}

	socialMedia, err := service.UpdateSocialMedia(c, request)
	if err != nil {
		response.ErrorResponse(
			c,
			http.StatusUnprocessableEntity,
			err.Error(),
		)
		log.Println("error update social media: ", err.Error())
		return
	}

	response.SuccessResponse(c, response.BaseResponse{
		Data:       socialMedia,
		Message:    "successfully updated social media",
		StatusCode: 200,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	idSocialMedia, err := strconv.Atoi(id)
	socialMedia, err := service.DeleteSocialMedia(c, idSocialMedia)
	if err != nil {
		response.ErrorResponse(
			c,
			http.StatusUnprocessableEntity,
			err.Error(),
		)
		log.Println("error delete social media by user: ", err.Error())
		return
	}

	response.SuccessResponse(c, response.BaseResponse{
		Data:       socialMedia,
		Message:    fmt.Sprintf("successfully deleted social media with id %s", id),
		StatusCode: 200,
	})
}

func GetSocialMediasByUser(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	}

	socialMedia, err := service.GetSocialMediasByUser(c, username)
	log.Print("sm: ", socialMedia)
	if err != nil {
		response.ErrorResponse(
			c,
			http.StatusUnprocessableEntity,
			err.Error(),
		)
		log.Println("error get social media by user: ", err.Error())
		return
	}

	response.SuccessResponse(c, response.BaseResponse{
		Data:       socialMedia,
		Message:    fmt.Sprintf("successfully get social media %s", username),
		StatusCode: 200,
	})
}

func GetDetailSocialMedia(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	idSocialMedia, err := strconv.Atoi(id)
	socialMedia, err := service.GetDetailSocialMedia(c, idSocialMedia)
	if err != nil {
		response.ErrorResponse(
			c,
			http.StatusUnprocessableEntity,
			err.Error(),
		)
		log.Println("error get social media by user: ", err.Error())
		return
	}

	response.SuccessResponse(c, response.BaseResponse{
		Data:       socialMedia,
		Message:    "successfully get social media",
		StatusCode: 200,
	})
}
