package app

import (
	"database/sql"
	"learning-database-migration-golang/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase() *sql.DB {
	// for execute database migration : migrate -database "mysql://root@tcp(localhost:3306)/learning_database_migration" -path db/migrations up
	// for create database migration : migrate create -ext sql -dir db/migrations create_table_first
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/learning_database_migration")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
