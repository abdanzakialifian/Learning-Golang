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

func TestSelectQuerySql(t *testing.T) {
	defer db.Close()

	ctx := context.Background()

	query := "select * from customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
	}

	defer rows.Close()
}
