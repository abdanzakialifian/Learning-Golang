package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) (T1, T2) {
	fmt.Println(param1)
	fmt.Println(param2)
	return param1, param2
}

func TestMultipleParameter(t *testing.T) {
	param1, param2 := MultipleParameter("Zaki", 100)
	assert.Equal(t, "Zaki", param1)
	assert.Equal(t, 100, param2)

	param3, param4 := MultipleParameter(100, "Zaki")
	assert.Equal(t, 100, param3)
	assert.Equal(t, "Zaki", param4)
}
