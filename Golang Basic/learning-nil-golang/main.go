package main

import "fmt"

func main() {
	data := NewMap("Zaki")

	if data == nil {
		fmt.Println("Empty Data")
	} else {
		fmt.Println(data["name"])
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		newMap := make(map[string]string)
		newMap["name"] = name
		return newMap
	}
}
