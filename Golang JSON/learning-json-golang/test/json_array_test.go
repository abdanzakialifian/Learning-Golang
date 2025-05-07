package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeJsonArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Abdan",
		MiddleName: "Zaki",
		LastName:   "Alifian",
		Age:        25,
		Married:    false,
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestDecodeJsonArray(t *testing.T) {
	jsonString := `{"FirstName":"Abdan","MiddleName":"Zaki","LastName":"Alifian","Age":25,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := new(Customer)
	if err := json.Unmarshal(jsonBytes, customer); err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
}

func TestEncodeJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Abdan",
		MiddleName: "Zaki",
		LastName:   "Alifian",
		Age:        25,
		Married:    false,
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
		Addresses: []Address{
			{
				Street:     "Jln. Dipayuda No. 10",
				Country:    "Indonesia",
				PostalCode: 12345,
			},
			{
				Street:     "Jln. Semampir No. 20",
				Country:    "Indonesia",
				PostalCode: 6789,
			},
		},
	}
	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestDecodeJsonArrayComplex(t *testing.T) {
	jsonString := `{"FirstName":"Abdan","MiddleName":"Zaki","LastName":"Alifian","Age":25,"Married":false,"Hobbies":["Gaming","Reading","Coding"],"Addresses":[{"Street":"Jln. Dipayuda No. 10","Country":"Indonesia","PostalCode":12345},{"Street":"Jln. Semampir No. 20","Country":"Indonesia","PostalCode":6789}]}`
	jsonBytes := []byte(jsonString)

	customer := new(Customer)
	if err := json.Unmarshal(jsonBytes, customer); err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}

func TestOnlyJsonArrayEncode(t *testing.T) {
	customer := Customer{
		Addresses: []Address{
			{
				Street:     "Jln. Dipayuda No. 10",
				Country:    "Indonesia",
				PostalCode: 12345,
			},
			{
				Street:     "Jln. Semampir No. 20",
				Country:    "Indonesia",
				PostalCode: 6789,
			},
		},
	}
	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestOnlyJsonArrayDecode(t *testing.T) {
	jsonString := `[{"Street":"Jln. Dipayuda No. 10","Country":"Indonesia","PostalCode":12345},{"Street":"Jln. Semampir No. 20","Country":"Indonesia","PostalCode":6789}]`
	jsonBytes := []byte(jsonString)

	addresses := new([]Address)
	if err := json.Unmarshal(jsonBytes, addresses); err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}
