package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min(100, 200))
	assert.Equal(t, int64(100), Min(int64(100), int64(100)))
	assert.Equal(t, float64(100), Min(float64(100), float64(200)))
	assert.Equal(t, "a", Min("a", "b"))
}
