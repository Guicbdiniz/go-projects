package user

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Guicbdiniz/go-projects/rest-api/api"
	"github.com/Guicbdiniz/go-projects/rest-api/db/queries"
	"github.com/Guicbdiniz/go-projects/rest-api/models"
	"github.com/Guicbdiniz/go-projects/rest-api/utils"

	_ "github.com/lib/pq"
)

func setUpTestDBForUserHandler(t *testing.T) *sql.DB {
	t.Log("set up test case")

	db, err := sql.Open("postgres", api.DatabaseTestingUrl)
	utils.CheckTestError(t, err, "Error captured while creating a db")

	err = queries.CreateUsersTable(db)
	utils.CheckTestError(t, err, "Error captured while creating users table")

	initialUser := models.User{
		ID:       0,
		Username: "guidi",
		Password: "senha",
	}
	err = initialUser.SaveUser(db)
	utils.CheckTestError(t, err, "Error captured while creating initial user")

	return db
}

func tearDownUserHandlerTestCase(t *testing.T, db *sql.DB) {
	t.Log("teardown test case")

	err := queries.DropUsersTable(db)
	utils.CheckTestError(t, err, "Error captured while droping user table")

	err = db.Close()
	utils.CheckTestError(t, err, "Error captured while closing db connection")
}

func TestUserHandler(t *testing.T) {
	db := setUpTestDBForUserHandler(t)
	defer tearDownUserHandlerTestCase(t, db)

	handler := CreateUserHandler(db)

	testGetAllUsers(t, db, handler)
	testCreateNewUser(t, db, handler)
}

func testGetAllUsers(t *testing.T, db *sql.DB, handler UserHandler) {
	request := httptest.NewRequest(http.MethodGet, "/user", nil)
	responseRecorder := httptest.NewRecorder()

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
}

func testCreateNewUser(t *testing.T, db *sql.DB, handler UserHandler) {
	requestBody := `{"username":"test_user","password":"password"}`

	request := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(requestBody))
	responseRecorder := httptest.NewRecorder()

	handler.ServeHTTP(responseRecorder, request)
	response := responseRecorder.Result()

	utils.AssertEqual(t, http.StatusCreated, response.StatusCode, "POST request to /user did not return the correct status")

	body, err := io.ReadAll(response.Body)
	utils.CheckTestError(t, err, "Error captured while reading a response")

	jsonBody, err := utils.UnmarshalJsonResponse[models.User](body)
	utils.CheckTestError(t, err, "Error captured while unsmarshiling users")

	utils.AssertEqual(t, "test_user", jsonBody.Data.Username, "GET request to /user did not return correct error text")
	utils.AssertEqual(t, "", jsonBody.ErrorText, "GET request to /user did not return correct error text")
}
