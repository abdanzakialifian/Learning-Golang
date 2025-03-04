package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Abdan Zaki Alifian", "ki"))
	fmt.Println(strings.Split("Abdan Zaki Alifian", " "))
	fmt.Println(strings.ToLower("Abdan Zaki Alifian"))
	fmt.Println(strings.ToUpper("Abdan Zaki Alifian"))
	fmt.Println(strings.Trim("          Abdan Zaki Alifian          ", " "))
	fmt.Println(strings.ReplaceAll("Abdan Zaki Alifian Zaki Alifian", "Zaki Alifian", ""))
}
