package main

import (
	"fmt"
	"reflect"
)

type Author struct {
	Name         string   `json:Name`
	Publications []string `json:Publications,omitempy`
}

func main() {
	chenglong := Author{Name: "ChengLong", Publications: []string{"QingHua", "BaiDa"}}
	fmt.Printf("%v", chenglong)
	fmt.Printf("%v+", chenglong)
	t := reflect.TypeOf(chenglong)
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		s, _ := t.FieldByName(name)
		fmt.Println(name, s.Tag)
	}
	fmt.Println("Main func done!")
}
