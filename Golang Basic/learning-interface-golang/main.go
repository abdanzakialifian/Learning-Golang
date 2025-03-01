package main

import "fmt"

func main() {
	person := Person{Name: "Zaki"}

	fmt.Println(person.GetName())

	sayHello(person)

	fmt.Println("===========================")

	animal := Animal{Name: "Kucing"}

	fmt.Println(animal.GetName())

	sayHello(animal)

	fmt.Println("===========================")

	fmt.Println(emptyInterface())
}

type HashName interface {
	GetName() string
}

func sayHello(value HashName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func emptyInterface() any {
	// return 1
	// return true
	return "Empty Interface (any)"
}
