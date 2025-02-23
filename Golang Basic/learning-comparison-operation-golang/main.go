package main

import "fmt"

func main() {
	name1 := "Zaki"
	name2 := "Zaki"

	result1 := name1 == name2
	result2 := name1 != name2

	fmt.Println(result1)
	fmt.Println(result2)

	number1 := 10
	number2 := 5

	result3 := number1 > number2
	result4 := number1 < number2

	fmt.Println(result3)
	fmt.Println(result4)
}
