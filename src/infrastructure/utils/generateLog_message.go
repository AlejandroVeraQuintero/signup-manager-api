package utils

import (
	"log"
	"strconv"
	"strings"

	message "github.com/AlejandroVeraQuintero/signup-manager-api/src/infrastructure/resources"
	"github.com/gin-gonic/gin"
)

const stringMethod = "{Method}"
const stringPath = "{Path}"
const stringStatusCode = "{StatusCode}"

func GenerateLogSuccessRequest(context *gin.Context) {
	log.Println(message.StringSeparator)
	log.Println(FormatMessageRequestLog(message.MessageSuccessRequest,
		context.Request.Method,
		context.Request.URL.Path, context.Writer.Status()))
	log.Println(message.StringSeparator)
}

func GenerateLogErrorRequest(context *gin.Context) {
	log.Println(message.StringSeparator)
	log.Println(FormatMessageRequestLog(message.MessageExceptionRequest,
		context.Request.Method,
		context.Request.URL.Path, context.Writer.Status()))
	log.Println(message.MessageDisplayError + context.Errors.String())
	log.Println(message.StringSeparator)
}

func FormatMessageRequestLog(message, method, path string, statusCode int) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(message, stringMethod, method), stringPath, path),
		stringStatusCode,
		strconv.Itoa(statusCode))
}
