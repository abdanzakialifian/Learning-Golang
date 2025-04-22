package test

import (
	"context"
	"fmt"
	"testing"
)

var db = GetConnection()

func TestInsertExecSql(t *testing.T) {
	defer db.Close()

	ctx := context.Background()

	query := "insert into customer values (\"P003\",\"Alifian\")"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestDeleteExecSql(t *testing.T) {
	defer db.Close()

	ctx := context.Background()

	query := "delete from customer where id = \"P001\""
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success delete customer data")
}

func TestUpdateExecSql(t *testing.T) {
	defer db.Close()

	ctx := context.Background()

	query := "update customer set name = \"Zaki\" where id = \"P002\""
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success update customer data")
}
