package test

import (
	"database/sql"
	"encoding/json"
	"go-restful-api/app"
	"go-restful-api/controller"
	"go-restful-api/helper"
	"go-restful-api/middleware"
	"go-restful-api/repository"
	"go-restful-api/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

// Function setup for connection to database test
func setupTestDB() *sql.DB {
	// (1) Open connection to database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_restful_golang_test")
	// (2) If error handle with helper
	helper.PanicErr(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Second)

	return db
}

// Function for handle router endpoint
func setupRouter() http.Handler {
	// (1) Run open connection db with function setupTestDB
	db := setupTestDB()
	// (2) Use validator
	validate := validator.New()

	// (3) Endpoint
	categoryRespository := repository.NewCategoriRepository()
	categoryService := service.NewCategoryService(categoryRespository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// (4) Use file router
	router := app.NewRouter(categoryController)

	// (5) Return router with handle middleware
	return middleware.NewAuthMiddleware(router)
}

// Function test for create category success
func TestCreateCategorySuccess(t *testing.T) {
	// (1) Use router
	router := setupRouter()

	// (2) Create request body payload
	requestBody := strings.NewReader(`{"name": "Gadget"}`)
	// (3) Create test request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	// (3) Added header content type
	request.Header.Add("Content-Type", "application/json")
	// (4) Added header authorize
	request.Header.Add("X-API-Key", "RAHASIA")

	// (5) Create new recorder for writer
	recorder := httptest.NewRecorder()

	// (6) Run test with send request
	router.ServeHTTP(recorder, request)

	// (7) Get result test and save to variable response
	response := recorder.Result()

	// (8) Read response body json
	body, _ := io.ReadAll(response.Body)
	// (9) Create variable responseBody with value map for response body
	var responseBody map[string]interface{}
	// (10) Decode json
	json.Unmarshal(body, &responseBody)

	// (11) Check response code must be 200
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	// (12) Check response status must be `OK`
	assert.Equal(t, "OK", responseBody["status"])
	// (13) Check response data must be gadget, and convert to type map
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}
