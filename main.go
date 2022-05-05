package main

import (
	"go-restful-api/app"
	"go-restful-api/controller"
	"go-restful-api/helper"
	"go-restful-api/middleware"
	"go-restful-api/repository"
	"go-restful-api/service"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// use db
	db := app.NewDB()
	// Use validator
	validate := validator.New()

	categoryRespository := repository.NewCategoriRepository()
	categoryService := service.NewCategoryService(categoryRespository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// Use file router
	router := app.NewRouter(categoryController)

	// Create server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	// Run server
	err := server.ListenAndServe()
	// If error handle with helper
	helper.PanicErr(err)

}
