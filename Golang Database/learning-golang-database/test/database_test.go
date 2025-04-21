package test

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnectionDatabase(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/learning_golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
