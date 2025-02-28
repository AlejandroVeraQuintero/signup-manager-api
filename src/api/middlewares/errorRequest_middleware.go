package middlewares

import (
	"net/http"

	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/utils"
	"github.com/gin-gonic/gin"
)

func LogErrorRequest() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		if context.Writer.Status() == http.StatusBadRequest {
			utils.CreateLogErrorRequest(context)
		}
	}
}
