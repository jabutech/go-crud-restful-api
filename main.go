package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jabutech/go-crud-restful-api/app"
	"github.com/jabutech/go-crud-restful-api/controller"
	"github.com/jabutech/go-crud-restful-api/helper"
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Create server
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// If no error, print message url run
	fmt.Println("App running at http://localhost:" + port)

	// Run server
	err := server.ListenAndServe()
	// If error handle with helper
	helper.PanicErr(err)

}
