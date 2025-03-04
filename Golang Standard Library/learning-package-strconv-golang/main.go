package main

import (
	"fmt"
	"strconv"
)

func main() {
	resultBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultBool)
	}

	fmt.Println("===============================")

	resultInt, err := strconv.Atoi("10000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	fmt.Println("===============================")

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	fmt.Println("===============================")

	resultString := strconv.Itoa(999)
	fmt.Println(resultString)

}
