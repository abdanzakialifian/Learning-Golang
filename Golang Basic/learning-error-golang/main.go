package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := division(100, 0)
	if err == nil {
		fmt.Println("Result", result)
	} else {
		fmt.Println("Error", err.Error())
	}
}

func division(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("division with zero")
	} else {
		return value / divider, nil
	}
}
