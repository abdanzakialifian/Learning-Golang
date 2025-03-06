package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	value := "Abdan Zaki Alifian"
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	fmt.Println("================================")

	bytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(bytes))
	}

	fmt.Println("================================")

	csvString := "abdan,zaki,alifian\n" + "budi,paratama,nugraha\n" + "joko,morro,diah"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

	fmt.Println("================================")

	writer := csv.NewWriter(os.Stdout)

	writer.Write([]string{"abdan", "zaki", "alifian"})
	writer.Write([]string{"budi", "pratama", "nugraha"})
	writer.Write([]string{"joko", "morro", "diah"})

	writer.Flush()
}
