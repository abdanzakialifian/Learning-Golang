package main

import "fmt"

func main() {
	var firstName string

	firstName = "Abdan"
	fmt.Println(firstName)

	firstName = "Zaki"
	fmt.Println(firstName)

	var lastName = "Alifian"
	fmt.Println(lastName)

	fullName := "Abdan Zaki Alifian"
	fmt.Println(fullName)

	fullName = "Abdan Alifian"
	fmt.Println(fullName)

	var (
		firstName1  = "Abdan"
		middleName1 = "Zaki"
		lastName1   = "Alifian"
	)
	fmt.Println(firstName1)
	fmt.Println(middleName1)
	fmt.Println(lastName1)
}
