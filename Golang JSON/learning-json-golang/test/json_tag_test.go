package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestEncodeJsonTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Apple Macbook Pro M1",
		ImageURL: "https://example.com/apple_macbook_pro_m1.png",
	}
	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestDecodeJsonTag(t *testing.T) {
	jsonString := `{"id":"P001","name":"Apple Macbook Pro M1","image_url":"https://example.com/apple_macbook_pro_m1.png"}`
	jsonBytes := []byte(jsonString)

	product := new(Product)

	if err := json.Unmarshal(jsonBytes, product); err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.ImageURL)
}
