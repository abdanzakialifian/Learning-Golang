package main

import "fmt"

func main() {
	type IdCardNumber string

	var idCardNumber1 IdCardNumber = "12345"
	var idCardNumber2 string = "678910"
	var idCardNumber3 IdCardNumber = IdCardNumber(idCardNumber2)

	fmt.Println(idCardNumber1)
	fmt.Println(idCardNumber3)
}
