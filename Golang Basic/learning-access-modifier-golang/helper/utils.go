package helper

import "fmt"

func AccessVersionWithSayGoodBye(name string) {
	result := sayGoodBye(name)
	fmt.Println(result)
	fmt.Println(version)
}
