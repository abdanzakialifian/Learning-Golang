package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	group := sync.WaitGroup{}
	var counter int64 = 0

	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	group.Wait()
	fmt.Println("Counter = ", counter)
}
