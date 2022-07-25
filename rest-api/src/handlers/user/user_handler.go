package user

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Guicbdiniz/go-projects/rest-api/models"
	"github.com/Guicbdiniz/go-projects/rest-api/utils"
)

type UserHandler struct {
	db *sql.DB
}

func CreateUserHandler(db *sql.DB) UserHandler {
	return UserHandler{
		db: db,
	}
}

func (h UserHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	default:
		utils.SendTextResponse(res, http.StatusMethodNotAllowed, "Invalid method")
		return
	case http.MethodGet:
		h.handleGetRequest(res, req)
		return
	case http.MethodPost:
		h.handlePostRequest(res, req)
		return
	}
}

func (h UserHandler) handleGetRequest(res http.ResponseWriter, req *http.Request) {
	requestParams := req.URL.Query()

	if requestParams.Has("username") {
		h.handleGetOneUserRequest(res, req, requestParams.Get("username"))
	} else {
		h.handleGetAllUsersRequest(res, req)
	}
}

func (h UserHandler) handleGetOneUserRequest(res http.ResponseWriter, req *http.Request, username string) {
	user, err := models.ReadUserByUsername(username, h.db)

	if err != nil {
		utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
		return
	}

	jsonResponse, err := utils.MarshalJsonResponse[models.User](user)

	if err != nil {
		utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
		return
	}

	utils.SendJsonResponse(res, http.StatusOK, jsonResponse)
}

func (h UserHandler) handleGetAllUsersRequest(res http.ResponseWriter, req *http.Request) {
	users, err := models.ReadAllUsers(h.db)

	if err != nil {
		utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
		return
	}

	jsonResponse, err := utils.MarshalJsonResponse[[]models.User](users)

	if err != nil {
		utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
		return
	}

	utils.SendJsonResponse(res, http.StatusOK, jsonResponse)
}

func (h UserHandler) handlePostRequest(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)

	if err != nil {
		utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
		return
	}

	var requestBody CreateUserRequestBody
	err = json.Unmarshal(body, &requestBody)

	if err != nil {
		utils.SendJsonErrorResponse(res, http.StatusBadRequest, err)
		return
	}

	user := models.User{Username: requestBody.Username, Password: requestBody.Password}
	err = user.SaveUser(h.db)

	if err != nil {
		utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
		return
	}

	jsonResponse, err := utils.MarshalJsonResponse[models.User](user)

	if err != nil {
		utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
		return
	}

	utils.SendJsonResponse(res, http.StatusCreated, jsonResponse)
}
