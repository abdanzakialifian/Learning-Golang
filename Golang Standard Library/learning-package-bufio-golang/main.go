package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("This is long string\nwith new line\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	fmt.Println("==========================")

	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("Hello world\n")
	writer.WriteString("Have a good study!\n")
	writer.Flush()
}
