package main

import (
	"go-restful-api/app"
	"go-restful-api/controller"
	"go-restful-api/exception"
	"go-restful-api/helper"
	"go-restful-api/repository"
	"go-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	// use db
	db := app.NewDB()
	// Use validator
	validate := validator.New()

	categoryRespository := repository.NewCategoriRepository()
	categoryService := service.NewCategoryService(categoryRespository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// Use http router
	router := httprouter.New()

	// Endpoint
	// Get all categories
	router.GET("/api/categories", categoryController.FindAll)
	// Get category by id
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	// Create new category
	router.POST("/api/categories", categoryController.Create)
	// Update category by id
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	// Delete category by id
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// Change PanicHandler to exception error hanlder
	router.PanicHandler = exception.ErrorHandler

	// Create server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	// Run server
	err := server.ListenAndServe()
	// If error handle with helper
	helper.PanicErr(err)

}
