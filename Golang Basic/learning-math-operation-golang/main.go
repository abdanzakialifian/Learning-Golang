package main

import "fmt"

func main() {
	a := 10
	b := 10
	c := 5
	d := 2
	e := a + b - c*d
	fmt.Println(e)

	i := 10
	i += 10
	fmt.Println(i)

	i += 5
	fmt.Println(i)

	i -= 10
	fmt.Println(i)

	j := 1
	j++
	fmt.Println(j)

	j++
	fmt.Println(j)

	j--
	fmt.Println(j)

	j--
	fmt.Println(j)
}
