package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name,omitempty"`
	SKU         string   `json:"sku"`
	Category    Category `json:"category"`
	Description string   `json:"-"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {
	p := Product{ID: 183, Name: "Iphone 14 Pro Max", SKU: "Iphone 14", Category: Category{ID: 3, Name: "APPLE"}}
	b, err := json.MarshalIndent(p, "", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	p2 := Product{ID: 39, SKU: "MeiZu 13"}
	b2, err := json.MarshalIndent(p2, "", "    ")
	fmt.Println(string(b2))

	fmt.Println("Main func done!")
}
