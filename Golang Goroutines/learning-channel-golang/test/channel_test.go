package test

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Abdan Zaki Alifian"
		fmt.Println("Complete send data to channel")
	}()

	data := <-channel
	println(data)

	time.Sleep(5 * time.Second)
}
