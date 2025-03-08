package main

import (
	"fmt"

	apphellogo "github.com/abdanzakialifian/app-hello-go"
	apphellogo2 "github.com/abdanzakialifian/app-hello-go/v2"
)

func main() {
	sayHello := apphellogo.SayHello()
	fmt.Println(sayHello)

	fmt.Println("========================")

	sayHello = apphellogo2.SayHello("Zaki")
	fmt.Println(sayHello)
}
