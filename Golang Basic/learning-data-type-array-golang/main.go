package main

import "fmt"

func main() {
	var names [3]string
	fmt.Println(names)

	names[0] = "Abdan"
	names[1] = "Zaki"
	names[2] = "Alifian"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(len(names))

	values := [3]int{
		70,
		80,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(len(values))

	values2 := [...]int{
		10,
		20,
		30,
		40,
		50,
	}
	fmt.Println(values2)
	fmt.Println(values2[0])
	fmt.Println(values2[1])
	fmt.Println(values2[2])
	fmt.Println(values2[3])
	fmt.Println(values2[4])
	fmt.Println(len(values2))
}
