package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Print[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestTypeParameter(t *testing.T) {
	resultString := Print("Zaki")
	assert.Equal(t, "Zaki", resultString)

	resultNumber := Print(100)
	assert.Equal(t, 100, resultNumber)
}
