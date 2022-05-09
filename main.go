package main

import (
	"net/http"

	"github.com/jabutech/go-crud-restful-api/app"
	"github.com/jabutech/go-crud-restful-api/controller"
	"github.com/jabutech/go-crud-restful-api/helper"
	"github.com/jabutech/go-crud-restful-api/middleware"
	"github.com/jabutech/go-crud-restful-api/repository"
	"github.com/jabutech/go-crud-restful-api/service"

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
