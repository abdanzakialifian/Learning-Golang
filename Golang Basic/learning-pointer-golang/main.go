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

	address4 := &address1 // copy by reference
	address4.City = "Tegal"
	fmt.Println(address1) // change
	fmt.Println(address4) // change

	fmt.Println("========================")

	address3 = &Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	}
	fmt.Println(address1) // do not change
	fmt.Println(address3) // change

	fmt.Println("========================")

	*address4 = Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	}
	address4.City = "Jakarta Selatan"
	fmt.Println(address1) // change because using asterisk operator (*)
	fmt.Println(address4) // change because using asterisk operator (*)

	fmt.Println("========================")

	address5 := new(Address)
	address6 := address5

	address6.Country = "Indonesia"

	fmt.Println(address5)
	fmt.Println(address6)

	fmt.Println("========================")

	address7 := Address{}
	changeCountry(&address7)
	fmt.Println(address7)

	fmt.Println("========================")

	zaki := Man{
		Name: "Zaki",
	}
	zaki.Married()
	fmt.Println(zaki.Name)
}

type Address struct {
	City     string
	Province string
	Country  string
}

func changeCountry(address *Address) {
	address.Country = "Indonesia"
}

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}
