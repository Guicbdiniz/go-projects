package utils

import (
	"fmt"
	"net/http"
)

func SendTextResponse(res http.ResponseWriter, httpCode int, message string) {
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(httpCode)
	res.Write([]byte(message))
}

func SendJsonResponse(res http.ResponseWriter, httpCode int, message []byte) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Content-Length", fmt.Sprintf("%d", len(message)))
	res.WriteHeader(httpCode)
	res.Write(message)
}

func SendJsonErrorResponse(res http.ResponseWriter, httpCode int, err error) {
	jsonErrorResponse, _ := MarshalJsonErrorResponse(err.Error())

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Content-Length", fmt.Sprintf("%d", len(jsonErrorResponse)))
	res.WriteHeader(httpCode)
	res.Write([]byte(jsonErrorResponse))
}
