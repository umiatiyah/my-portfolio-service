package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
}

func Response(context *gin.Context, statusCode int, data interface{}) {
	context.JSON(statusCode, data)
}

func SuccessResponse(context *gin.Context, data interface{}) {
	Response(context, http.StatusOK, data)
}

func ErrorResponse(context *gin.Context, statusCode int, message string) {
	Response(context, statusCode, BaseResponse{StatusCode: statusCode, Message: message})
}
