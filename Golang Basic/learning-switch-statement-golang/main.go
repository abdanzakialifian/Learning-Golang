package main

import "fmt"

func main() {
	name := "Zaki"

	switch name {
	case "Abdan":
		fmt.Println("Hello Abdan")
	case "Zaki":
		fmt.Println("Hello Zaki")
	case "Alifian":
		fmt.Println("Hello Alifian")
	default:
		fmt.Println("Hi, can I get to know you?")
	}

	fmt.Println("=========================")

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name is too long")
	case false:
		fmt.Println("Name is correct")
	}

	fmt.Println("=========================")

	name = "Alifian"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Name is too long")
	case length > 5:
		fmt.Println("Name is quite long")
	default:
		fmt.Println("Name is correct")
	}
}
