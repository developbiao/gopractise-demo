package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]*Student

func main() {
	list = make(map[string]*Student)
	student := Student{"LiJun"}
	list["lijun"] = &student
	list["lijun"].Name = "KaoLa"
	fmt.Println(list["lijun"].Name)
	fmt.Println("Main func done!")
}
