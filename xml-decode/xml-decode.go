package main

import (
	"encoding/xml"
	"fmt"
)

type Cat struct {
	Name string `xml:"name"`
	Age  uint   `xml:"age"`
}

type MyXML struct {
	Cat `xml:"cat"`
}

func main() {
	myXMLStr := []byte(`<cat>
    <name>Timi</name>
    <age>5</age>
    </cat>
    `)
	c := MyXML{}
	err := xml.Unmarshal(myXMLStr, &c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Cat.Name)
	fmt.Println(c.Cat.Age)
	fmt.Println("Main func done!")
}
