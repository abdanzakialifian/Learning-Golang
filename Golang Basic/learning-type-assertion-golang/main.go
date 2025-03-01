package main

import "fmt"

func main() {
	result := Random()
	resultString := result.(string)
	fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt)

	fmt.Println("==============================")

	switch value := result.(type) {
	case string:
		fmt.Println("String data type", value)
	case int:
		fmt.Println("Int data type", value)
	default:
		fmt.Println("Unknown", value)
	}
}

func Random() any {
	return "OK"
}
