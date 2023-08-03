package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data := []byte("Hello Goalng!\n")
	fmt.Println(data)
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1")
	fmt.Println(string(read1))
}
