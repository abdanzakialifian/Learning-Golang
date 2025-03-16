package test

import (
	"fmt"
	"strconv"
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

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		for i := range 10 {
			channel <- "Loop " + strconv.Itoa(i+1)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Get data :", data)
	}

	fmt.Println("Done")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1 : ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2 : ", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1 : ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2 : ", data)
			counter++
		default:
			fmt.Println("Waiting data...")
		}

		if counter == 2 {
			break
		}
	}
}
