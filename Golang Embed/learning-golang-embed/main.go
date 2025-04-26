package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed test/file/version.txt
var version string

//go:embed test/file/golang_logo.png
var logo []byte

//go:embed test/file/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("golang_logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, err := path.ReadDir("test/file")
	if err != nil {
		panic(err)
	}
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, err := path.ReadFile("test/file/" + entry.Name())
			if err != nil {
				panic(err)
			}
			fmt.Println(string(content))
		}
	}
}
