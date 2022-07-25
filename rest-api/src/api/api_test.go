package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Guicbdiniz/go-projects/rest-api/configs"
	"github.com/Guicbdiniz/go-projects/rest-api/utils"

	_ "github.com/lib/pq"
)

func TestCreateAPI(t *testing.T) {
	api, err := CreateAPI(configs.DatabaseTestingUrl)

	utils.CheckTestError(t, err, "Error captured while creating API")

	defer api.CloseDBConnection()

	request := httptest.NewRequest("GET", "/ping", nil)
	responseRecorder := httptest.NewRecorder()

	api.ServeHTTP(responseRecorder, request)

	response := responseRecorder.Result()

	utils.AssertEqual(t, http.StatusOK, response.StatusCode, "Ping request did not return the correct status code")

}
