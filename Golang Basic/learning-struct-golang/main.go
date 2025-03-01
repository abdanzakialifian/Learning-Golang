package main

import "fmt"

func main() {
	var zaki Customer

	fmt.Println(zaki)

	fmt.Println("===========================")

	zaki.Name = "Abdan Zaki Alifian"
	zaki.Address = "Indonesia"
	zaki.Age = 25

	fmt.Println(zaki)

	fmt.Println("===========================")

	fmt.Println(zaki.Name)
	fmt.Println(zaki.Address)
	fmt.Println(zaki.Age)

	fmt.Println("===========================")

	abdan := Customer{
		Name:    "Abdan",
		Address: "Indonesia",
		Age:     25,
	}

	fmt.Println(abdan)

	fmt.Println("===========================")

	alifian := Customer{"Alifian", "Indonesia", 25}

	fmt.Println(alifian)
}

type Customer struct {
	Name    string
	Address string
	Age     int
}
