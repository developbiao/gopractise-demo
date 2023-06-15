package main

import "fmt"

const (
	a = iota
	b = iota
)
const (
	name = "chengdu"
	c    = iota
	d    = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(name)
	fmt.Println(c)
	fmt.Println(d)
}
