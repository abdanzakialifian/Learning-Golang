package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		How to test arguments os in golang? by running the following code :
		- go run main.go Abdan Zaki Alifian
		- go run main.go "Abdan Zaki Alifian"
	*/
	args := os.Args

	fmt.Println(args)

	fmt.Println("================================")

	for _, arg := range args {
		fmt.Println(arg)
	}

	fmt.Println("================================")

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
