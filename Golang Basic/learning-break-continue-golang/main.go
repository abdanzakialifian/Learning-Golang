package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Loop to", i)
	}

	fmt.Println("========================")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Loop to", i)
	}
}
