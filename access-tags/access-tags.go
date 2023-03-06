package main

import (
	"fmt"
	"reflect"
)

type Product struct {
	ID          uint64 `json:"id" test:"Yo Yooo"`
	Name        string `json:"name,omitempty"`
	Description string `json:"-"`
}

func main() {
	p := Product{ID: 3}
	t := reflect.TypeOf(p)
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("Field Name: %s \n", t.Field(i).Name)
		fmt.Printf("Field Tag: %s \n", t.Field(i).Tag)
	}

	// Lookup tag exists
	if tagValue, ok := t.Field(0).Tag.Lookup("test"); ok {
		fmt.Println(tagValue)
	} else {
		fmt.Println("No tag 'test'")
	}

	fmt.Println("Main func done!")
}
