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

	fmt.Println("================================")

	err = saveData("test")

	if err != nil {
		switch error := err.(type) {
		case *validationError:
			fmt.Println("validation error:", error.Error())
		case *notFoundError:
			fmt.Println("not found error:", error.Error())
		default:
			fmt.Println("unknown error:", error.Error())
		}
	} else {
		fmt.Println("Success")
	}
}

func division(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("division with zero")
	} else {
		return value / divider, nil
	}
}

func saveData(id string) error {
	if id == "" {
		return &validationError{
			Message: "validation error",
		}
	}

	if id != "zaki" {
		return &notFoundError{
			Message: "data not found",
		}
	}

	return nil
}

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}
