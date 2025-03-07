package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	fmt.Println(path.Dir("hello/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))

	fmt.Println("============================================")

	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.IsAbs("hello//world.go"))
	fmt.Println(filepath.IsLocal("hello/world.go"))
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}
