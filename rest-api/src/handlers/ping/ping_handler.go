package ping

import (
	"net/http"

	"github.com/Guicbdiniz/go-projects/rest-api/utils"
)

func Ping(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		utils.SendTextResponse(res, http.StatusMethodNotAllowed, "Invalid method")
	}

	utils.SendTextResponse(res, http.StatusOK, "pong")
}
