package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number2 interface {
	// ~ is type approximation
	~int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func Min2[T Number2](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin2(t *testing.T) {
	assert.Equal(t, Age(100), Min2(Age(100), Age(200)))
}
