package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEcoder(t *testing.T) {
	writer, _ := os.Create("customer-out.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Rangga",
		MiddleName: "Surya",
		LastName:   "Prayoga",
	}

	encoder.Encode(customer)
	fmt.Println(customer)
}