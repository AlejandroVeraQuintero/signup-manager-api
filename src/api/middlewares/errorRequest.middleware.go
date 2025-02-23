package middlewares

import (
	"log"
	"net/http"

	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/message"
	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/utils"
	"github.com/gin-gonic/gin"
)

func LogErrorRequest() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		if context.Writer.Status() == http.StatusBadRequest {
			log.Println(message.StringSeparator)
			log.Println(utils.FormatMessageRequestLog(message.MessageExceptionRequest,
				context.Request.Method,
				context.Request.URL.Path, context.Writer.Status()))
			log.Println(message.MessageDisplayError + context.Errors.String())
			log.Println(message.StringSeparator)
		}
	}
}
