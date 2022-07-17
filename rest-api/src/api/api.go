package api

import (
	"net/http"

	"github.com/Guicbdiniz/go-projects/rest-api/handlers/ping"
)

type API struct {
	serveMux *http.ServeMux
}

// Register the handler function for the given pattern in the ServeMux.
func (api *API) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	api.serveMux.HandleFunc(pattern, handler)
}

// Register the handler for the given pattern in the ServeMux.
func (api *API) Handle(pattern string, handler http.Handler) {
	api.serveMux.Handle(pattern, handler)
}

// Dispatch the request to the handler whose pattern most closely matches the request URL.
func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.serveMux.ServeHTTP(w, r)
}

func CreateAPI() *API {
	api := API{}
	api.serveMux = http.NewServeMux()

	api.HandleFunc("/ping", ping.Ping)

	return &api
}
