package test

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
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

var group = sync.WaitGroup{}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine :", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine :", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter :", n)
		if n == 10 {
			break
		}
	}

	cancel()

	group.Wait()

	fmt.Println("Total Goroutine :", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine :", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine :", runtime.NumGoroutine())

	for n := range destination {
		println("Counter :", n)
	}

	group.Wait()

	fmt.Println("Total Goroutine :", runtime.NumGoroutine())
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	group.Add(1)
	go func() {
		defer close(destination)
		defer group.Done()
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}
