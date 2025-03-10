package test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can't run on MAC OS")
	}

	fmt.Println("This code not execute (skip unit test)")
}
