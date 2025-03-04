package main

import (
	"flag"
	"fmt"
)

func main() {
	/*
		How to test flag in golang? by running the following code :
		- go run main.go -username=zaki -password=rahasia -host=123.321.46.7 -port=5505
		- go run main.go -username=zaki -password="rahasia sekali" -host=123.321.46.7 -port=5505
	*/
	username := flag.String("username", "root", "database username")
	password := flag.String("password", "root", "database password")
	host := flag.String("host", "localhost", "database host")
	port := flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)
}
