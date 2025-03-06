package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Jhon", "Paul", "George", "Ringgo"}
	values := []int{100, 95, 90, 80}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Zaki"))
	fmt.Println(slices.Contains(names, "Paul"))
	fmt.Println(slices.Index(names, "Zaki"))
	fmt.Println(slices.Index(names, "Paul"))
}
