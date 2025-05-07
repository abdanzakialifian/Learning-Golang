package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeJsonMap(t *testing.T) {
	product := map[string]any{
		"id":    "P002",
		"name":  "HP Pavillion",
		"price": 14000000,
		"components": []string{
			"SSD",
			"RAM",
			"VGA",
		},
	}
	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestDecodeJsonMap(t *testing.T) {
	jsonString := `{"id":"P001","name":"Apple Macbook Pro M1","price":20000000}`
	jsonBytes := []byte(jsonString)

	result := map[string]any{}

	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}
