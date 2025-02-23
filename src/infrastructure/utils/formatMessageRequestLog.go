package utils

import (
	"strconv"
	"strings"
)

const stringMethod = "{Method}"
const stringPath = "{Path}"
const stringStatusCode = "{StatusCode}"

func FormatMessageRequestLog(message, method, path string, statusCode int) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(message, stringMethod, method), stringPath, path),
		stringStatusCode,
		strconv.Itoa(statusCode))
}
