package main

import (
	"encoding/xml"
	"fmt"
)

type Price struct {
	Text     string `xml:",chardata"`
	Currency string `xml:"currency,attr"`
}

type Product struct {
	Price   Price  `xml:"price"`
	Name    string `xml:"name"`
	Comment string `xml:",comment"`
}

func main() {
	p := Product{Name: "Iphone 13", Price: Price{Text: "8000", Currency: "CNY"}, Comment: "一个好的果子"}
	b, err := xml.MarshalIndent(p, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	fmt.Println("Main func done!")
}
