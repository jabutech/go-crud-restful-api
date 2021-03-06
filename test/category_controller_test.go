package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jabutech/go-crud-restful-api/app"
	"github.com/jabutech/go-crud-restful-api/controller"
	"github.com/jabutech/go-crud-restful-api/helper"
	"github.com/jabutech/go-crud-restful-api/middleware"
	"github.com/jabutech/go-crud-restful-api/model/domain"
	"github.com/jabutech/go-crud-restful-api/repository"
	"github.com/jabutech/go-crud-restful-api/service"
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

// Function for handle router endpoint with parameter connetion to db
func setupRouter(db *sql.DB) http.Handler {
	// (1) Use validator
	validate := validator.New()

	// (2) Endpoint
	categoryRespository := repository.NewCategoriRepository()
	categoryService := service.NewCategoryService(categoryRespository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// (3) Use file router
	router := app.NewRouter(categoryController)

	// (4) Return router with handle middleware
	return middleware.NewAuthMiddleware(router)
}

// Function for truncate table category
func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

// Function test for create category success
func TestCreateCategorySuccess(t *testing.T) {
	// (1) Use connetion to db
	db := setupTestDB()
	// (2) Run truncate table category before test
	truncateCategory(db)
	// (3) Use router
	router := setupRouter(db)

	// (4) Create request body payload
	requestBody := strings.NewReader(`{"name": "Gadget"}`)
	// (5) Create test request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	// (6) Added header content type
	request.Header.Add("Content-Type", "application/json")
	// (7) Added header authorize
	request.Header.Add("X-API-Key", "RAHASIA")

	// (8) Create new recorder for writer
	recorder := httptest.NewRecorder()

	// (9) Run test with send request
	router.ServeHTTP(recorder, request)

	// (10) Get result test and save to variable response
	response := recorder.Result()

	// (11) Read response body json
	body, _ := io.ReadAll(response.Body)
	// (12) Create variable responseBody with value map for response body
	var responseBody map[string]interface{}
	// (13) Decode json
	json.Unmarshal(body, &responseBody)

	// (14) Response status code must be 200 (success)
	assert.Equal(t, 200, response.StatusCode)
	// (15) Check response body code must be 200 (success)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	// (16) Check response status must be `OK`
	assert.Equal(t, "OK", responseBody["status"])
	// (17) Check response data must be gadget, and convert to type map
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

// Function test for create category failed
func TestCreateCategoryFailed(t *testing.T) {
	// (1) Use connetion to db
	db := setupTestDB()
	// (2) Run truncate table category before test
	truncateCategory(db)
	// (3) Use router
	router := setupRouter(db)

	// (4) Create request body payload
	requestBody := strings.NewReader(`{"name": ""}`)
	// (5) Create test request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	// (6) Added header content type
	request.Header.Add("Content-Type", "application/json")
	// (7) Added header authorize
	request.Header.Add("X-API-Key", "RAHASIA")

	// (8) Create new recorder for writer
	recorder := httptest.NewRecorder()

	// (9) Run test with send request
	router.ServeHTTP(recorder, request)

	// (10) Get result test and save to variable response
	response := recorder.Result()

	// (11) Read response body json
	body, _ := io.ReadAll(response.Body)
	// (12) Create variable responseBody with value map for response body
	var responseBody map[string]interface{}
	// (13) Decode json
	json.Unmarshal(body, &responseBody)

	// (14) Response status code must be 400 (bad request)
	assert.Equal(t, 400, response.StatusCode)
	// (15) Check response body code must be 400 (bad request)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	// (16) Check response status must be `OK`
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

// Function test for update category success
func TestUpdateCategorySuccess(t *testing.T) {
	// (1) Use connetion to db
	db := setupTestDB()
	// (2) Run truncate table category before test
	truncateCategory(db)

	// (3) Create new data for sample update
	// (3.1) Create database transactional
	tx, _ := db.Begin()
	// (3.2) Use repository
	categoryRepository := repository.NewCategoriRepository()
	// (3.3) Create new category
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	// (3.4) Commit transaction
	tx.Commit()

	// (4) Use router
	router := setupRouter(db)

	// (4) Create request body payload update
	requestBody := strings.NewReader(`{"name": "T SHIRT"}`)
	// (5) Create test request update with id
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	// (6) Added header content type
	request.Header.Add("Content-Type", "application/json")
	// (7) Added header authorize
	request.Header.Add("X-API-Key", "RAHASIA")

	// (8) Create new recorder for writer
	recorder := httptest.NewRecorder()

	// (9) Run test with send request
	router.ServeHTTP(recorder, request)

	// (10) Get result test and save to variable response
	response := recorder.Result()

	// (11) Read response body json
	body, _ := io.ReadAll(response.Body)
	// (12) Create variable responseBody with value map for response body
	var responseBody map[string]interface{}
	// (13) Decode json
	json.Unmarshal(body, &responseBody)

	// (14) Response status code must be 200 (success)
	assert.Equal(t, 200, response.StatusCode)
	// (15) Check response body code must be 200 (success)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	// (16) Check response status must be `OK`
	assert.Equal(t, "OK", responseBody["status"])
	// (17) Check response data id must be same as the id above, and convert to type map
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	// (17) Check response data name must be gadget, and convert to type map
	assert.Equal(t, "T SHIRT", responseBody["data"].(map[string]interface{})["name"])
}

// Function test for update category failed
func TestUpdateCategoryFailed(t *testing.T) {
	// (1) Use connetion to db
	db := setupTestDB()
	// (2) Run truncate table category before test
	truncateCategory(db)

	// (3) Create new data for sample update
	// (3.1) Create database transactional
	tx, _ := db.Begin()
	// (3.2) Use repository
	categoryRepository := repository.NewCategoriRepository()
	// (3.3) Create new category
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	// (3.4) Commit transaction
	tx.Commit()

	// (5) Use router
	router := setupRouter(db)

	// (6) Create request body payload update
	requestBody := strings.NewReader(`{"name": ""}`)
	// (7) Create test request update with id
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	// (8) Added header content type
	request.Header.Add("Content-Type", "application/json")
	// (9) Added header authorize
	request.Header.Add("X-API-Key", "RAHASIA")

	// (10) Create new recorder for writer
	recorder := httptest.NewRecorder()

	// (11) Run test with send request
	router.ServeHTTP(recorder, request)

	// (12) Get result test and save to variable response
	response := recorder.Result()

	// (13) Read response body json
	body, _ := io.ReadAll(response.Body)
	// (14) Create variable responseBody with value map for response body
	var responseBody map[string]interface{}
	// (15) Decode json
	json.Unmarshal(body, &responseBody)

	// (16) Response status code must be 400 (BAD REQUEST)
	assert.Equal(t, 400, response.StatusCode)
	// (17) Check response body code must be 400 (BAD REQUEST)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	// (18) Check response status must be `OK`
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

// Function test for get category success
func TestGETCategorySuccess(t *testing.T) {
	// (1) Use connetion to db
	db := setupTestDB()
	// (2) Run truncate table category before test
	truncateCategory(db)

	// (3) Create new data for sample update
	// (3.1) Create database transactional
	tx, _ := db.Begin()
	// (3.2) Use repository
	categoryRepository := repository.NewCategoriRepository()
	// (3.3) Create new category
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	// (3.4) Commit transaction
	tx.Commit()

	// (4) Use router
	router := setupRouter(db)

	// (5) Create test request update with id
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	// (7) Added header authorize
	request.Header.Add("X-API-Key", "RAHASIA")

	// (8) Create new recorder for writer
	recorder := httptest.NewRecorder()

	// (9) Run test with send request
	router.ServeHTTP(recorder, request)

	// (10) Get result test and save to variable response
	response := recorder.Result()

	// (11) Read response body json
	body, _ := io.ReadAll(response.Body)
	// (12) Create variable responseBody with value map for response body
	var responseBody map[string]interface{}
	// (13) Decode json
	json.Unmarshal(body, &responseBody)

	// (14) Response status code must be 200 (success)
	assert.Equal(t, 200, response.StatusCode)
	// (15) Check response body code must be 200 (success)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	// (16) Check response status must be `OK`
	assert.Equal(t, "OK", responseBody["status"])
	// (17) Check response data id must be same as the id above, and convert to type map
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	// (17) Check response data name must be gadget, and convert to type map
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

// Function test for get category failed
func TestGETCategoryFailed(t *testing.T) {
	// (1) Use connetion to db
	db := setupTestDB()
	// (2) Run truncate table category before test
	truncateCategory(db)
	// (3) Use router
	router := setupRouter(db)

	// (5) Create test request update with id
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/404", nil)
	// (7) Added header authorize
	request.Header.Add("X-API-Key", "RAHASIA")

	// (8) Create new recorder for writer
	recorder := httptest.NewRecorder()

	// (9) Run test with send request
	router.ServeHTTP(recorder, request)

	// (10) Get result test and save to variable response
	response := recorder.Result()

	// (11) Read response body json
	body, _ := io.ReadAll(response.Body)
	// (12) Create variable responseBody with value map for response body
	var responseBody map[string]interface{}
	// (13) Decode json
	json.Unmarshal(body, &responseBody)

	// (14) Response status code must be 200 (Not Found)
	assert.Equal(t, 404, response.StatusCode)
	// (15) Check response body code must be 404 (Not Found)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	// (16) Check response status must be `NOT FOUND`
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

// Function test for delete category success
func TestDeleteCategorySuccess(t *testing.T) {
	// (1) Use connetion to db
	db := setupTestDB()
	// (2) Run truncate table category before test
	truncateCategory(db)

	// (3) Create new data for sample update
	// (3.1) Create database transactional
	tx, _ := db.Begin()
	// (3.2) Use repository
	categoryRepository := repository.NewCategoriRepository()
	// (3.3) Create new category
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	// (3.4) Commit transaction
	tx.Commit()

	// (4) Use router
	router := setupRouter(db)

	// (6) Create test request update with id
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	// (7) Added header content type
	request.Header.Add("Content-Type", "application/json")
	// (8) Added header authorize
	request.Header.Add("X-API-Key", "RAHASIA")

	// (9) Create new recorder for writer
	recorder := httptest.NewRecorder()

	// (10) Run test with send request
	router.ServeHTTP(recorder, request)

	// (11) Get result test and save to variable response
	response := recorder.Result()

	// (12) Read response body json
	body, _ := io.ReadAll(response.Body)
	// (13) Create variable responseBody with value map for response body
	var responseBody map[string]interface{}
	// (14) Decode json
	json.Unmarshal(body, &responseBody)

	// (15) Response status code must be 200 (success)
	assert.Equal(t, 200, response.StatusCode)
	// (16) Check response body code must be 200 (success)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	// (17) Check response status must be `OK`
	assert.Equal(t, "OK", responseBody["status"])
}

// Function test for delete category failed
func TestDeleteCategoryfailed(t *testing.T) {
	// (1) Use connetion to db
	db := setupTestDB()
	// (2) Run truncate table category before test
	truncateCategory(db)

	// (3) Use router
	router := setupRouter(db)

	// (4) Create test request delete with id
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/404", nil)
	// (6) Added header content type
	request.Header.Add("Content-Type", "application/json")
	// (7) Added header authorize
	request.Header.Add("X-API-Key", "RAHASIA")

	// (8) Create new recorder for writer
	recorder := httptest.NewRecorder()

	// (9) Run test with send request
	router.ServeHTTP(recorder, request)

	// (10) Get result test and save to variable response
	response := recorder.Result()

	// (11) Read response body json
	body, _ := io.ReadAll(response.Body)
	// (12) Create variable responseBody with value map for response body
	var responseBody map[string]interface{}
	// (13) Decode json
	json.Unmarshal(body, &responseBody)

	// (14) Response status code must be 404 (NOT FOUND)
	assert.Equal(t, 404, response.StatusCode)
	// (15) Check response body code must be 404 (NOT FOUND)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	// (16) Check response status must be `OK`
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

// Function test for get list category success
func TestGetListCategorySuccess(t *testing.T) {
	// (1) Use connetion to db
	db := setupTestDB()
	// (2) Run truncate table category before test
	truncateCategory(db)

	// (3) Create new data for sample update
	// (3.1) Create database transactional
	tx, _ := db.Begin()
	// (3.2) Use repository
	categoryRepository := repository.NewCategoriRepository()
	// (3.3) Create new category
	category1 := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	// (3.4) Create new category
	category2 := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "T SHIRT",
	})
	// (3.5) Commit transaction
	tx.Commit()

	// (4) Use router
	router := setupRouter(db)

	// (5) Create test request get all categories
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	// (6) Added header content type
	request.Header.Add("Content-Type", "application/json")
	// (7) Added header authorize
	request.Header.Add("X-API-Key", "RAHASIA")

	// (8) Create new recorder for writer
	recorder := httptest.NewRecorder()

	// (9) Run test with send request
	router.ServeHTTP(recorder, request)

	// (10) Get result test and save to variable response
	response := recorder.Result()

	// (11) Read response body json
	body, _ := io.ReadAll(response.Body)
	// (12) Create variable responseBody with value map for response body
	var responseBody map[string]interface{}
	// (13) Decode json
	json.Unmarshal(body, &responseBody)

	// (14) Response status code must be 200 (success)
	assert.Equal(t, 200, response.StatusCode)
	// (15) Check response body code must be 200 (success)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	// (16) Check response status must be `OK`
	assert.Equal(t, "OK", responseBody["status"])
	// (17) Convertion response body to interface
	var categories = responseBody["data"].([]interface{})
	// (18) Convertion to map
	categoryResponse1 := categories[0].(map[string]interface{})
	categoryResponse2 := categories[1].(map[string]interface{})
	// (19) Check id response 1 must be the same with category id 1
	assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64)))
	// (20) Check name response 1 must be the same with category name 1
	assert.Equal(t, category1.Name, categoryResponse1["name"])
	// (21) Check id response 2 must be the same with category id 2
	assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64)))
	// (22) Check name response 2 must be the same with category name 2
	assert.Equal(t, category2.Name, categoryResponse2["name"])

}

// Function test for unauthorized
func TestUnauthorized(t *testing.T) {
	// (1) Use connetion to db
	db := setupTestDB()
	// (2) Run truncate table category before test
	truncateCategory(db)

	// (4) Use router
	router := setupRouter(db)

	// (5) Create test request get all categories
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	// (6) Added header content type
	request.Header.Add("Content-Type", "application/json")
	// (7) Added header wrong authorized API Key
	request.Header.Add("X-API-Key", "SALAH")

	// (8) Create new recorder for writer
	recorder := httptest.NewRecorder()

	// (9) Run test with send request
	router.ServeHTTP(recorder, request)

	// (10) Get result test and save to variable response
	response := recorder.Result()

	// (11) Read response body json
	body, _ := io.ReadAll(response.Body)
	// (12) Create variable responseBody with value map for response body
	var responseBody map[string]interface{}
	// (13) Decode json
	json.Unmarshal(body, &responseBody)

	// (14) Response status code must be 401 (UNAUTHORIZED)
	assert.Equal(t, 401, response.StatusCode)
	// (15) Check response body code must be 401 (UNAUTHORIZED)
	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	// (16) Check response status must be `OK`
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])

}
