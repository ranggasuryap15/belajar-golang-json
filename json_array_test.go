package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName": "Rangga", "MiddleName": "Surya", "LastName": "Prayoga", "Age": 21, "Married": false, "Hobbies": ["Coding", "Coding with You"]}`
	customer := Customer{}
	json.Unmarshal([]byte(jsonString), &customer)
	fmt.Println(customer)
}

func TestJsonArrayEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "Rangga",
		MiddleName: "Surya",
		LastName:   "Prayoga",
		Age:        21,
		Married:    false,
		Hobbies:    []string{"Coding", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}