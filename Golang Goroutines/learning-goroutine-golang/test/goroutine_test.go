package test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestManyGoroutine(t *testing.T) {
	for i := range 100000 {
		go DisplayNumber(i + 1)
	}

	time.Sleep(5 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}
