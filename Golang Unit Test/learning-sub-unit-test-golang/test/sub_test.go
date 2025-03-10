package test

import (
	"learning-sub-unit-test-golang/helper"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Test1", func(t *testing.T) {
		result := helper.SayHello("Zaki")
		require.Equal(t, "Hello Zaki", result, "Result must be 'Hello Zaki'")
	})
	t.Run("Test2", func(t *testing.T) {
		result := helper.SayHello("Abdan")
		require.Equal(t, "Hello Abdan", result, "Result must be 'Hello Abdan'")
	})
}
