package test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
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

func TestInsertCustomer(t *testing.T) {
	defer db.Close()

	ctx := context.Background()

	query := "insert into customer (id, name, email, balance, rating, birth_date, married) values ('P003','Alifian','alifian@gmail.com',3000000,70.0,'1999-07-7',false)"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
}

func TestSelectCustomer(t *testing.T) {
	defer db.Close()

	ctx := context.Background()

	query := "select id, name, email, balance, rating, birth_date, married, created_at from customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id string
		var name string
		var email sql.NullString
		var balance int
		var rating float32
		var birthDate sql.NullTime
		var married bool
		var createdAt time.Time

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		if email.Valid {
			fmt.Println("Email :", email.String)
		}
		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date :", birthDate.Time)
		}
		fmt.Println("Married :", married)
		fmt.Println("Created At :", createdAt)

		fmt.Println("=============================")
	}

	defer rows.Close()
}
