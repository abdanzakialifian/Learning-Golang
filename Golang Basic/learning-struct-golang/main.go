package main

import "fmt"

func main() {
	var zaki Customer

	fmt.Println(zaki)

	fmt.Println("===========================")

	zaki.Name = "Zaki"
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

	fmt.Println("===========================")

	abdan.sayHello("Steven")
	zaki.sayHello("Steven")
	alifian.sayHello("Steven")
}

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", customer.Name, "my name is", name)
}
