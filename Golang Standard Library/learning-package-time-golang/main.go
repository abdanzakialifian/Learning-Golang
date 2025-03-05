package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	fmt.Println("============================================")

	utc := time.Date(2000, time.August, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	fmt.Println("============================================")

	formatter := "2006-01-02 15:04:05"
	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
		fmt.Println(valueTime.Year())
		fmt.Println(valueTime.Month())
		fmt.Println(valueTime.Day())
		fmt.Println(valueTime.Hour())
	}

	fmt.Println("============================================")

	duration1 := 100 * time.Second
	duration2 := 10 * time.Millisecond
	duration3 := duration1 - duration2

	fmt.Println(duration1)
	fmt.Println(duration2)
	fmt.Println(duration3)
	fmt.Printf("Duration : %d\n", duration3)
}
