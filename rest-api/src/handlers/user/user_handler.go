package user

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Guicbdiniz/go-projects/rest-api/models"
	"github.com/Guicbdiniz/go-projects/rest-api/utils"
)

type UserHandler struct {
	db *sql.DB
}

func (h UserHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		users, err := models.ReadAllUsers(h.db)

		if err != nil {
			log.Println(err.Error())
			utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
		}

		jsonResponse, err := utils.MarshalJsonResponse[[]models.User](users)

		if err != nil {
			utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
		}

		utils.SendJsonResponse(res, http.StatusOK, jsonResponse)
		return
	default:
		utils.SendTextResponse(res, http.StatusMethodNotAllowed, "Invalid method")
		return
	}
}

func CreateUserHandler(db *sql.DB) UserHandler {
	return UserHandler{
		db: db,
	}
}
