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

	fmt.Println("===================")

	sayHelloWithFilter("Zaki", spamFilter)

	sayHelloWithTypeDeclaration("Alifian", spamFilter)

	filtered := spamFilter
	sayHelloWithFilter("Anjing", filtered)

	fmt.Println("===================")

	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	registerUser("Zaki", blacklist)
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
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

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

type Filter func(string) string

func sayHelloWithTypeDeclaration(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

type Blacklist func(string) bool

func registerUser(name string, isBlacklist Blacklist) {
	if isBlacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
