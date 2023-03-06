package main

import (
	"encoding/xml"
	"fmt"
)

type Product struct {
	ID          int64    `xml:"id"`
	Name        string   `xml:"ownName,omitempty"`
	SKU         string   `xml:"sku"`
	Category    Category `xml:"category"`
	Description string   `xml:"-"`
}

type Category struct {
	ID   int64  `xml:"id"`
	Name string `xml:"name"`
}

func main() {
	p := Product{ID: 183, Name: "Iphone 14 Pro Max", SKU: "Iphone 14", Category: Category{ID: 3, Name: "APPLE"}}
	b, err := xml.MarshalIndent(p, "", "    ")
	if err != nil {
		panic(err)
	}
	xmlWithHeader := xml.Header + string(b)
	fmt.Println(xmlWithHeader)
	fmt.Println("Main func done!")
}
