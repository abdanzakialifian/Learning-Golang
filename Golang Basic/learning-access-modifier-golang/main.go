package main

import (
	"fmt"
	"learning-access-modifier-golang/helper"
)

func main() {
	result := helper.SayHello("Zaki")
	fmt.Println(result)
	fmt.Println(helper.Application)

	fmt.Println("=================================")

	helper.AccessVersionWithSayGoodBye("Zaki")
}
