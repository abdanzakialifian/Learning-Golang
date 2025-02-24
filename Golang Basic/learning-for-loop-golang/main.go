package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Loop to", counter)
		counter++
	}
	fmt.Println("Done")

	fmt.Println("============================")

	for couting := 1; couting <= 10; couting++ {
		fmt.Println("Loop to", couting)
	}
	fmt.Println("Done")

	fmt.Println("============================")

	names := []string{"Abdan", "Zaki", "Alifian"}
	for index, name := range names {
		fmt.Println("Index", index+1, name)
	}
}
