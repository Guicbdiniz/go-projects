package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Guicbdiniz/go-projects/gin-api/internal/utils"
)

func TestCreateAPI(t *testing.T) {
	api, err := CreateAPI()

	utils.CheckTestError(t, err, "CreateAPI returned an error")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	api.ServeHTTP(w, req)

	utils.AssertEqual(t, http.StatusOK, w.Code, "Ping route did not return http.StatusOK")
}
