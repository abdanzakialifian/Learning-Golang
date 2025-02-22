package main

import "fmt"

func main() {
	const value16 = 32769
	const value32 = int32(value16)
	const value64 = int64(value32)

	fmt.Println(value16)
	fmt.Println(value32)
	fmt.Println(value64)

	name := "Zaki"
	byteCharacter := name[0]
	byteString := string(byteCharacter)

	fmt.Println(byteCharacter)
	fmt.Println(byteString)
}
