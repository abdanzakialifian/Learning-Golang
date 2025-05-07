package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode int
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Abdan",
		MiddleName: "Zaki",
		LastName:   "Alifian",
		Age:        25,
		Married:    true,
	}
	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
