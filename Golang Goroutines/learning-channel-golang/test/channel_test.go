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

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Abdan Zaki Alifian"
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Abdan Zaki alifian"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Abdan"
	channel <- "Zaki"
	channel <- "Alifian"

	fmt.Println("Capacity Channel :", cap(channel))
	fmt.Println("Size Channel :", len(channel))

	fmt.Println("============== Get Channel ==============")

	fmt.Println(<-channel)
	fmt.Println("Size Channel :", len(channel))
	fmt.Println(<-channel)
	fmt.Println("Size Channel :", len(channel))
	fmt.Println(<-channel)
	fmt.Println("Size Channel :", len(channel))

	fmt.Println("============== Done ==============")
}
