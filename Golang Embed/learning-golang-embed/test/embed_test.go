package test

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed file/version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed file/golang_logo.png
var logo []byte

func TestByte(t *testing.T) {
	fmt.Println(logo)
	err := ioutil.WriteFile("golang_logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed file/a.txt
//go:embed file/b.txt
//go:embed file/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, err := files.ReadFile("file/a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(a))

	b, err := files.ReadFile("file/b.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	c, err := files.ReadFile("file/c.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(c))
}

//go:embed file/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, err := path.ReadDir("file")
	if err != nil {
		panic(err)
	}
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, err := path.ReadFile("file/" + entry.Name())
			if err != nil {
				panic(err)
			}
			fmt.Println(string(content))
		}
	}
}
