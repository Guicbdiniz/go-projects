package handlers

import (
	"net/http"

	"github.com/Guicbdiniz/go-projects/rest-api/api"
)

type BaseHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, api *api.API)
}
