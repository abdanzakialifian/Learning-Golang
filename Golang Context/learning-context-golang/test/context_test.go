package test

import (
	"context"
	"fmt"
	"testing"
)

type contextKey string

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, contextKey("b"), "B")
	contextC := context.WithValue(contextA, contextKey("c"), "C")

	contextD := context.WithValue(contextB, contextKey("d"), "D")
	contextE := context.WithValue(contextB, contextKey("e"), "E")

	contextF := context.WithValue(contextC, contextKey("f"), "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println("================= Get Key =================")

	fmt.Println(contextF.Value(contextKey("f")))
	fmt.Println(contextF.Value(contextKey("c")))
	fmt.Println(contextF.Value(contextKey("b")))
}
