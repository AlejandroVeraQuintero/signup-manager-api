package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func HandleResponse[T any](context *gin.Context, data T, err error) {
	if err != nil {
		SendError(context, http.StatusBadRequest, err)
	} else {
		SendResponse(context, http.StatusOK, data)
	}
}

func SendResponse[T any](context *gin.Context, statusCode int, data T) {
	response := APIResponse[T]{
		Status: statusCode,
		Data:   data,
	}
	context.IndentedJSON(http.StatusOK, response)
}

func SendError(context *gin.Context, statusCode int, err error) {
	response := APIResponse[any]{
		Status:  statusCode,
		Message: err.Error(),
	}
	context.IndentedJSON(statusCode, response)
}
