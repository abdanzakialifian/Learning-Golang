package main

import "fmt"

func main() {
	names := [...]string{
		"Abdan",
		"Zaki",
		"Alifian",
		"Steven",
		"Khandar",
		"Brokoli",
	}

	slice1 := names[3:5]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	fmt.Println("=================================")

	slice2 := names[:3]
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	fmt.Println("=================================")

	slice3 := names[2:]
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	fmt.Println("=================================")

	slice4 := names[:]
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	fmt.Println("=================================")

	days := [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)
	daySlice1[0] = "New Saturday"
	daySlice1[1] = "New Sunday"
	fmt.Println(daySlice1)
	fmt.Println(days)

	fmt.Println("=================================")

	daySlice2 := append(daySlice1, "New Holiday")
	daySlice2[0] = "Old Saturday"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	fmt.Println("=================================")

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Abdan"
	newSlice[1] = "Zaki"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fmt.Println("=================================")

	newSlice2 := append(newSlice, "Alifian")
	newSlice2[1] = "Steven"
	fmt.Println(newSlice)
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	fmt.Println("=================================")

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	fmt.Println("=================================")

	thisArray := [...]int{1, 2, 3, 4, 5}
	thisSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(thisArray)
	fmt.Println(thisSlice)
}
