package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("Rangga")
	logJson(true)
	logJson(15)
	logJson([]string{"a", "b", "c", "d"})
}