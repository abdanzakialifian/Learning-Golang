package test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, err := os.Open("response/customer.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(reader)
	customer := new(Customer)
	decoder.Decode(customer)

	fmt.Println(customer)
}
