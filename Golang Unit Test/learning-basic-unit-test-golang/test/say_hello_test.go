package test

import (
	"learning-basic-unit-test-golang/helper"
	"testing"
)

func TestFirstSayHello(t *testing.T) {
	sayHello := helper.SayHello("Zaki")
	if sayHello != "Hello Zaki" {
		panic("Result is not a `Hello Zaki`")
	}
}

func TestSecondSayHello(t *testing.T) {
	sayHello := helper.SayHello("Abdan")
	if sayHello != "Hello Abdan" {
		panic("Result is not a `Hello Abdan`")
	}
}
