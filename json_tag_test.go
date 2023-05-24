package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	ID       string `json:"id`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product {
		ID: "P0001",
		Name: "Apple Mac Book Pro",
		ImageURL: "http://example.com/image.png",
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"id": "P0001", "name": "Apple Macbook Pro", "IMAGE_URL": "http://example.com/images.png"}`
	jsonBytes := []byte(jsonString)

	product := Product{}

	json.Unmarshal(jsonBytes, &product)
	fmt.Println(product)
}