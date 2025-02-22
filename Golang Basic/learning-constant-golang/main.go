package main

import "fmt"

func main() {
	const firstName = "Abdan"
	const middleName = "Zaki"
	const lastName = "Alifian"

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

	const (
		firstName1  = "Abdan 1"
		middleName1 = "Zaki 1"
		lastName1   = "Alifian 1"
	)

	fmt.Println(firstName1)
	fmt.Println(middleName1)
	fmt.Println(lastName1)
}
