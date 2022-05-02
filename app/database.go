package app

import (
	"database/sql"
	"go-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
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
