package middlewares

import (
	"log"
	"net/http"

	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/message"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/utils"
	"github.com/gin-gonic/gin"
)

func LogSuccessRequest() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		if context.Writer.Status() == http.StatusOK {
			log.Println(message.StringSeparator)
			log.Println(utils.FormatMessageRequestLog(message.MessageSuccessRequest,
				context.Request.Method,
				context.Request.URL.Path, context.Writer.Status()))
			log.Println(message.StringSeparator)
		}
	}
}
