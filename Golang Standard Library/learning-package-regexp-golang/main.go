package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`z[a-z]*i`)

	fmt.Println(regex.MatchString("zaki"))
	fmt.Println(regex.MatchString("zAKi"))
	fmt.Println(regex.MatchString("zii"))

	fmt.Println("================================")

	fmt.Println(regex.FindAllString("zaki zAKi ZAKI zii zoe zulham", 10))
}
