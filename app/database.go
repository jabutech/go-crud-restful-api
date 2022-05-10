package app

import (
	"database/sql"
	"os"
	"time"

	"github.com/jabutech/go-crud-restful-api/helper"
	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {
	// Load file .env
	godotenv.Load(".env")
	// Get variable DATABSE_URL from .env file
	dbUrl := os.Getenv("DATABASE_URL")

	// (1) Open connection to database
	db, err := sql.Open("mysql", dbUrl)
	// (2) If error handle with helper
	helper.PanicErr(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Second)

	return db
}
