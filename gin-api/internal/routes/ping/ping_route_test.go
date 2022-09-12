package ping

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Guicbdiniz/go-projects/gin-api/internal/utils"
	"github.com/gin-gonic/gin"
)

func TestPingRoute(t *testing.T) {
	router := gin.New()

	err := AddPingRoute(router)

	utils.CheckTestError(t, err, "AddPingRoute returned an error")

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(recorder, request)

	utils.AssertEqual(t, http.StatusOK, recorder.Code, "Ping route did not returned StatusOK")
	utils.AssertEqual(t, "pong", recorder.Body.String(), "Ping route did not returned 'pong' as body")
}
