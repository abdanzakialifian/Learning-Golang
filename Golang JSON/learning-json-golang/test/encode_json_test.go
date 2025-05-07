package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeJson(t *testing.T) {
	logJson("Zaki")
	logJson(29)
	logJson(true)
	logJson([]string{"Abdan", "Zaki", "Alifian"})
}

func logJson(data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
