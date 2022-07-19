package ping

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	request := httptest.NewRequest("POST", "/ping", nil)
	responseRecorder := httptest.NewRecorder()

	handler := PingHandler{}

	handler.ServeHTTP(responseRecorder, request)

	response := responseRecorder.Result()

	if response.StatusCode == http.StatusOK {
		t.Fatalf("Invalid response status code for a POST request. Expected %d, got %d",
			http.StatusMethodNotAllowed, response.StatusCode)
	}

	request = httptest.NewRequest("GET", "/ping", nil)
	responseRecorder = httptest.NewRecorder()

	handler.ServeHTTP(responseRecorder, request)

	response = responseRecorder.Result()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("Invalid response status code for a GET request. Expected %d, got %d",
			http.StatusOK, response.StatusCode)
	}

}
