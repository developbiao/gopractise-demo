package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	if err := encoder.Encode(data); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(filename, buffer.Bytes(), 0600); err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	decoder := gob.NewDecoder(buffer)
	if err := decoder.Decode(data); err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Id: 3, Author: "老树", Content: "老树画画"}
	store(post, "post3")
	var postRead Post
	load(&postRead, "post3")
	fmt.Println(postRead)
}
