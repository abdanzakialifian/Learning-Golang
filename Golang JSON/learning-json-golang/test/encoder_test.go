package test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, err := os.Create("response/customer_ouput.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(writer)
	customer := Customer{
		FirstName:  "Abdan",
		MiddleName: "Zaki",
		LastName:   "Alifian",
	}
	err = encoder.Encode(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}
