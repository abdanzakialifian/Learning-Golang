package main

import "fmt"

func main() {
	sayHello()

	fmt.Println("===================")

	sayHelloTo("Abdan", "Zaki", "Alifian")

	fmt.Println("===================")

	result := getHello("Zaki")
	fmt.Println(result)

	fmt.Println("===================")

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	firstName, _ = getFullName()
	fmt.Println(firstName)

	fmt.Println("===================")

	a, b, c := getCompleteName()
	fmt.Println(a, b, c)

	fmt.Println("===================")

	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 50))

	fmt.Println("===================")

	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
	fmt.Println(sumAll(10, 10, 10, 10))

	fmt.Println("===================")

	goodbye := getGoodBye
	fmt.Println(goodbye("Zaki"))
}

func sayHello() {
	fmt.Println("Hello Golang!")
}

func sayHelloTo(firstName string, middleName string, lastName string) {
	fmt.Println("Hello", firstName, middleName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Abdan", "Zaki"
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Abdan"
	middleName = "Zaki"
	lastName = "Alifian"
	return firstName, middleName, lastName
}

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func getGoodBye(name string) string {
	return "Good bye " + name
}
