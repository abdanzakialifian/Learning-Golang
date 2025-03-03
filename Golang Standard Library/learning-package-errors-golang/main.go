package main

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("validation error")
	ErrNotFound   = errors.New("not found error")
)

func main() {
	CheckId("zaki")
}

func CheckId(id string) {
	err := GetById(id)

	if err == nil {
		return
	}

	if errors.Is(err, ErrValidation) {
		fmt.Println(err.Error())
	} else if errors.Is(err, ErrNotFound) {
		fmt.Println(err.Error())
	} else {
		fmt.Println("unknown error")
	}
}

func GetById(id string) error {
	if id == "" {
		return ErrValidation
	} else if id != "zaki" {
		return ErrNotFound
	} else {
		return nil
	}
}
