package api

import (
	"net/http"

	"github.com/Guicbdiniz/go-projects/rest-api/handlers/ping"
)

func CreateAPI() *http.ServeMux {
	api := http.NewServeMux()

	api.HandleFunc("/ping", ping.Ping)

	return api
}
