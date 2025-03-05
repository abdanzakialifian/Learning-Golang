package main

import (
	"fmt"
	"sort"
)

func main() {
	users := []User{
		{
			Name: "Zaki",
			Age:  25,
		},
		{
			Name: "Abdan",
			Age:  22,
		},
		{
			Name: "Alifian",
			Age:  26,
		},
		{
			Name: "Stev",
			Age:  23,
		},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)

	for i := range users {
		fmt.Println("i :", i)
		for j := 0; j <= i; j++ {
			fmt.Println("j :", j)
		}
	}
}

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (us UserSlice) Len() int {
	return len(us)
}

func (us UserSlice) Less(i, j int) bool {
	return us[i].Age < us[j].Age
}

func (us UserSlice) Swap(i, j int) {
	us[i], us[j] = us[j], us[i]
}
