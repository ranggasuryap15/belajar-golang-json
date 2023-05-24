package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName": "John", "MiddleName": "Kocack", "LastName": "Doe", "Age": 30, "Married": true}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	json.Unmarshal(jsonBytes, &customer)
	fmt.Println(customer)
}