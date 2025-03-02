package main

import "fmt"

func main() {
	address1 := Address{
		City:     "Banjarnegara",
		Province: "Jawa Tengah",
		Country:  "Indonesia",
	}
	address2 := address1 // copy by value
	address2.City = "Purwokerto"
	fmt.Println(address1) // do not change
	fmt.Println(address2) // change

	fmt.Println("========================")

	address3 := &address1 // copy by reference
	address3.City = "Purbalingga"
	fmt.Println(address1) // change
	fmt.Println(address3) // change
}

type Address struct {
	City     string
	Province string
	Country  string
}
