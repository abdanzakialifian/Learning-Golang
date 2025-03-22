package test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}

	for range 100 {
		group.Add(1)
		go func() {
			defer group.Done()
			time.Sleep(3 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU :", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread :", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine :", totalGoroutine)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}

	for range 100 {
		group.Add(1)
		go func() {
			defer group.Done()
			time.Sleep(3 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU :", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread :", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine :", totalGoroutine)

	group.Wait()
}
