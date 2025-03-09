package test

import (
	"fmt"
	"learning-assertion-unit-test-golang/helper"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequireSayHello(t *testing.T) {
	sayHello := helper.SayHello("Zaki")
	require.Equal(t, "Hello Zaki", sayHello, "Result must be 'Hello Zaki'")
	fmt.Println("This code not execute (fail now)")
}

func TestAssertionSayHello(t *testing.T) {
	sayHello := helper.SayHello("Zaki")
	assert.Equal(t, "Hello Zaki", sayHello, "Result must be 'Hello Zaki'")
	fmt.Println("This code will be execute (fail)")
}
