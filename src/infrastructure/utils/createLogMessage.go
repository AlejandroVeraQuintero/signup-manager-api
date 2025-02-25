package utils

import (
	"log"

	"github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/message"
	"github.com/gin-gonic/gin"
)

func CreateLogSuccessRequest(context *gin.Context) {
	log.Println(message.StringSeparator)
	log.Println(FormatMessageRequestLog(message.MessageSuccessRequest,
		context.Request.Method,
		context.Request.URL.Path, context.Writer.Status()))
	log.Println(message.StringSeparator)
}

func CreateLogErrorRequest(context *gin.Context) {
	log.Println(message.StringSeparator)
	log.Println(FormatMessageRequestLog(message.MessageExceptionRequest,
		context.Request.Method,
		context.Request.URL.Path, context.Writer.Status()))
	log.Println(message.MessageDisplayError + context.Errors.String())
	log.Println(message.StringSeparator)
}
