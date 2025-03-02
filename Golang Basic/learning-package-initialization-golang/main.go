package main

import (
	"fmt"
	"learning-package-initialization-golang/database"
	_ "learning-package-initialization-golang/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
