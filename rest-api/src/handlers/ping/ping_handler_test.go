package ping

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Guicbdiniz/go-projects/rest-api/utils"
)

func TestPing(t *testing.T) {
	request := httptest.NewRequest("POST", "/ping", nil)
	responseRecorder := httptest.NewRecorder()

	handler := PingHandler{}

	handler.ServeHTTP(responseRecorder, request)

	response := responseRecorder.Result()

	utils.AssertEqual(t, http.StatusMethodNotAllowed, response.StatusCode, "Invalid response status code for a POST request")

	request = httptest.NewRequest("GET", "/ping", nil)
	responseRecorder = httptest.NewRecorder()

	handler.ServeHTTP(responseRecorder, request)

	response = responseRecorder.Result()

	utils.AssertEqual(t, http.StatusOK, response.StatusCode, "Invalid response status code for a GET request")

}
