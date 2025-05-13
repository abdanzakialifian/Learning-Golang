package app

import (
	"database/sql"
	"learning-restful-api-golang/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/learning_restful_api")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
