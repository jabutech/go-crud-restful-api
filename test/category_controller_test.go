package test

import (
	"database/sql"
	"go-restful-api/app"
	"go-restful-api/controller"
	"go-restful-api/helper"
	"go-restful-api/middleware"
	"go-restful-api/repository"
	"go-restful-api/service"
	"net/http"
	"time"

	"github.com/go-playground/validator"
)

// Function setup for connection to database test
func setupTestDB() *sql.DB {
	// (1) Open connection to database
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_restful_golang")
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
