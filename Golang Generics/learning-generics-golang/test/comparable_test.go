package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestComparable(t *testing.T) {
	assert.True(t, IsSame("Zaki", "Zaki"))
	assert.False(t, IsSame("Zaki", "zaki"))
	assert.True(t, IsSame(100, 100))
	assert.False(t, IsSame(100, 10))
}
