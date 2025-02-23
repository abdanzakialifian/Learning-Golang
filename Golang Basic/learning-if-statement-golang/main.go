package main

import "fmt"

func main() {
	name := "Zaki"

	if name == "Zaki" {
		fmt.Println("Hello", name)
	}

	fmt.Println("====================")

	name = "Abdan"
	if name == "Zaki" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Hi, can I get to know you?")
	}

	fmt.Println("====================")

	name = "Alifian"
	if name == "Abdan" {
		fmt.Println("Hello Abdan")
	} else if name == "Zaki" {
		fmt.Println("Hello Zaki")
	} else if name == "Alifian" {
		fmt.Println("Hello Alifian")
	} else {
		fmt.Println("Hi, can I get to know you?")
	}

	fmt.Println("====================")

	if length := len(name); length > 5 {
		fmt.Println("Name is too long")
	} else {
		fmt.Println("Name is correct")
	}
}
