package main

import (
	"fmt"
	"reflect"
)

func main() {
	readFiled(Sample{"Zaki"})

	fmt.Println("===========================")

	readFiled(Person{
		Name:    "Zaki",
		Address: "Banjarnegara",
		Email:   "test@gmail.com",
	})

	fmt.Println("===========================")

	person := Person{
		Name:    "Zaki",
		Address: "Banjarnegara",
		Email:   "test@gmail.com",
	}
	fmt.Println(isValid(person))
}

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readFiled(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := range valueType.NumField() {
		valueField := valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
		fmt.Println(valueField.Tag.Get("required"))
		fmt.Println(valueField.Tag.Get("max"))
	}
}

func isValid(value any) (result bool) {
	result = true
	reflectType := reflect.TypeOf(value)
	for i := range reflectType.NumField() {
		field := reflectType.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if !result {
				return result
			}
		}
	}
	return result
}
