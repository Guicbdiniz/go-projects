package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAPI(t *testing.T) {
	api := CreateAPI()

	request := httptest.NewRequest("GET", "/ping", nil)
	responseRecorder := httptest.NewRecorder()

	api.ServeHTTP(responseRecorder, request)

	response := responseRecorder.Result()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("Ping request did not return the correct status code. Expected %d, got %d",
			http.StatusAccepted, response.StatusCode)
	}

}
