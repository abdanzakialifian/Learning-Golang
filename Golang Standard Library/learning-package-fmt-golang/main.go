package main

import "fmt"

func main() {
	firstName := "Abdan"
	middleName := "Zaki"
	lastName := "Alifian"

	fmt.Println("Hello '", firstName, middleName, lastName, "'")
	fmt.Printf("Hello '%s %s %s'\n", firstName, middleName, lastName)
}
