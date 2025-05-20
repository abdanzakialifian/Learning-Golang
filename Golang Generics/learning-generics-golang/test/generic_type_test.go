package test

import (
	"fmt"
	"reflect"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	fmt.Println("Type:", reflect.TypeOf(bag))
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBag(t *testing.T) {
	names := Bag[string]{"Abdan", "Zaki", "Alifian"}
	PrintBag(names)

	numbers := Bag[int]{1, 2, 3, 4, 5}
	PrintBag(numbers)
}
