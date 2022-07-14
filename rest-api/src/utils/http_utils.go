package utils

import "net/http"

func SendTextResponse(res http.ResponseWriter, httpCode int, message string) {
	res.Header().Set("Content-type", "text/plain")
	res.WriteHeader(httpCode)
	res.Write([]byte(message))
}
