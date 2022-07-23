package user

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Guicbdiniz/go-projects/rest-api/db/queries"

	"github.com/Guicbdiniz/go-projects/rest-api/models"

	"github.com/Guicbdiniz/go-projects/rest-api/api"
	"github.com/Guicbdiniz/go-projects/rest-api/utils"

	_ "github.com/lib/pq"
)

func TestUserHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "/user", nil)
	responseRecorder := httptest.NewRecorder()

	db, err := sql.Open("postgres", api.DatabaseTestingUrl)
	utils.CheckTestError(t, err, "Error captured while creating a db")
	defer db.Close()

	err = queries.CreateUsersTable(db)
	utils.CheckTestError(t, err, "Error captured while creating users table")

	initialUser := models.User{
		ID:       0,
		Username: "guidi",
		Password: "senha",
	}
	err = initialUser.SaveUser(db)
	utils.CheckTestError(t, err, "Error captured while creating initial user")

	handler := CreateUserHandler(db)
	handler.ServeHTTP(responseRecorder, request)
	response := responseRecorder.Result()

	utils.AssertEqual(t, http.StatusOK, response.StatusCode, "GET request to /user did not return the correct status")

	body, err := io.ReadAll(response.Body)
	utils.CheckTestError(t, err, "Error captured while reading a response")

	jsonBody, err := utils.UnmarshalJsonResponse[[]models.User](body)
	utils.CheckTestError(t, err, "Error captured while unsmarshiling users")

	utils.AssertEqual(t, 1, len(jsonBody.Data), "GET request to /user did not return correct body")
	utils.AssertEqual(t, jsonBody.ErrorText, "", "GET request to /user did not return correct error text")
	utils.AssertEqual(t, jsonBody.Data[0].Username, "guidi", "GET request to /user did not return correct body")

	queries.DropUsersTable(db)
}
