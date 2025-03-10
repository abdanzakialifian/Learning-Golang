package test

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	// Before
	fmt.Println("Before Unit Test")
	m.Run()
	// After
	fmt.Println("After Unit Test")
}
