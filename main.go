package main

import (
	"net/http"
	"os"

	"github.com/jabutech/go-crud-restful-api/app"
	"github.com/jabutech/go-crud-restful-api/controller"
	"github.com/jabutech/go-crud-restful-api/helper"
	"github.com/jabutech/go-crud-restful-api/middleware"
	"github.com/jabutech/go-crud-restful-api/repository"
	"github.com/jabutech/go-crud-restful-api/service"
	"github.com/joho/godotenv"

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

	// Load file .env
	godotenv.Load(".env")

	// Get variable from env file
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8000"
	}
	addr := "localhost:" + appPort

	// Create server
	server := http.Server{
		Addr:    addr,
		Handler: middleware.NewAuthMiddleware(router),
	}

	// Run server
	err = server.ListenAndServe()
	// If error handle with helper
	helper.PanicErr(err)

}
