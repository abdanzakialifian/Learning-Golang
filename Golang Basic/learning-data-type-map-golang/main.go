package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Zaki",
		"address": "Banjarnegara",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(len(person))

	fmt.Println("=================================")

	book := make(map[string]string)
	book["title"] = "Golang Book"
	book["author"] = "Abdan Zaki Alifian"
	book["ups"] = "Wrong"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
