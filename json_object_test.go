package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "John",
		MiddleName: "Doe",
		LastName:   "Doe",
	}
	jsonString, err := json.Marshal(customer)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(jsonString))
}