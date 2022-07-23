package api

import (
	"database/sql"
	"net/http"

	"github.com/Guicbdiniz/go-projects/rest-api/handlers/ping"
)

const DatabaseTestingUrl = "postgresql://api_tester:password@localhost:5432/test_api"

type API struct {
	serveMux *http.ServeMux
	db       *sql.DB
}

// Register the handler for the given pattern in the ServeMux.
func (api *API) Handle(pattern string, handler http.Handler) {
	api.serveMux.Handle(pattern, handler)
}

// Dispatch the request to the handler whose pattern most closely matches the request URL.
func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.serveMux.ServeHTTP(w, r)
}

func (api *API) CloseDBConnection() error {
	return api.db.Close()
}

func CreateAPI(postgreUrl string) (*API, error) {
	db, err := sql.Open("postgres", postgreUrl)

	if err != nil {
		return &API{}, err
	}

	api := API{}
	api.serveMux = http.NewServeMux()
	api.db = db

	api.Handle("/ping", ping.CreatePingHandler(db))

	return &api, nil
}
