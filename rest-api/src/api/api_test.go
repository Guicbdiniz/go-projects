package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"
)

const databaseTestingUrl = "postgresql://api_tester:password@localhost:5432/test_api"

func TestCreateAPI(t *testing.T) {
	api, err := CreateAPI(databaseTestingUrl)

	if err != nil {
		t.Fatalf("Error captured while creating API: %s", err.Error())
	}

	defer api.CloseDBConnection()

	request := httptest.NewRequest("GET", "/ping", nil)
	responseRecorder := httptest.NewRecorder()

	api.ServeHTTP(responseRecorder, request)

	response := responseRecorder.Result()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("Ping request did not return the correct status code. Expected %d, got %d",
			http.StatusAccepted, response.StatusCode)
	}

}
