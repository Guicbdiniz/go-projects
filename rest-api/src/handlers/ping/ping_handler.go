package ping

import (
	"database/sql"
	"net/http"

	"github.com/Guicbdiniz/go-projects/rest-api/utils"
)

type PingHandler struct {
	db *sql.DB
}

func (h PingHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		utils.SendTextResponse(res, http.StatusMethodNotAllowed, "Invalid method")
	}

	utils.SendTextResponse(res, http.StatusOK, "pong")
}

func CreatePingHandler(db *sql.DB) PingHandler {
	h := PingHandler{}
	h.db = db
	return h
}
