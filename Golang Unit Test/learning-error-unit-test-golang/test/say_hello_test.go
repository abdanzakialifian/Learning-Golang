package test

import (
	"fmt"
	"learning-error-unit-test-golang/helper"
	"testing"
)

func TestFailSayHello(t *testing.T) {
	sayHello := helper.SayHello("Zaki")
	if sayHello != "Hello Zaki" {
		t.Fail()
	}

	fmt.Println("This code will be execute because handling error using fail")
}

func TestFailNowSayHello(t *testing.T) {
	sayHello := helper.SayHello("Abdan")
	if sayHello != "Hello Abdan" {
		t.FailNow()
	}

	fmt.Println("This code not execute because handling error using fail now")
}

func TestErrorSayHello(t *testing.T) {
	sayHello := helper.SayHello("Zaki")
	if sayHello != "Hello Zaki" {
		t.Error("Result must be 'Hello Zaki'")
	}

	fmt.Println("This code will be execute because handling error using error (fail)")
}

func TestFatalSayHello(t *testing.T) {
	sayHello := helper.SayHello("Abdan")
	if sayHello != "Hello Abdan" {
		t.Fatal("Result must be 'Hello Abdan'")
	}

	fmt.Println("This code not execute because handling error using fatal (fail now)")
}
