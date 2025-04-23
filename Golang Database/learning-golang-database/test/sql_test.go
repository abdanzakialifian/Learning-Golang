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

func TestSqlInjection(t *testing.T) {
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	query := "select username from user where username = '" + username + "' and password = '" + password + "' limit 1"
	fmt.Println(query)
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
	defer rows.Close()
}

func TestSqlIjectionSafe(t *testing.T) {
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	query := "select username from user where username = ? and password = ? limit 1"
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}

	defer rows.Close()
}

func TestExecSqlParameter(t *testing.T) {
	defer db.Close()

	ctx := context.Background()

	username := "zaki"
	password := "admin123"

	query := "insert into user (username,password) values (?,?)"
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}

func TestLastInsertId(t *testing.T) {
	defer db.Close()

	ctx := context.Background()

	email := "zaki@gmail.com"
	comment := "Test comment"

	query := "insert into comments(email, comment) values(?,?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", insertId)
}
