package main

import "fmt"

type Person struct {
	age int
}

func (p Person) howOld() int {
	return p.age
}

func (p *Person) growUp() int {
	p.age += 1
	return p.age
}

func main() {
	// caoqi value type
	caoqi := Person{age: 32}
	fmt.Println(caoqi.howOld())

	// caoqi grow up
	fmt.Println(caoqi.growUp())

	// pointer type
	nana := &Person{age: 16}
	fmt.Println(nana.howOld())
	fmt.Println(nana.growUp())
}
