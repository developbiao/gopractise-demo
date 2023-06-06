package main

import "fmt"

func main() {
	x := 300
	var any interface{} = x
	fmt.Println(any)
	var biao coder
	biao = Gopher{"Golang"}
	biao.code()
	biao.debug()

	var c coder
	fmt.Println(c == nil)
	fmt.Printf("c: %T, %v\n", c, c)

	var g *Gopher
	fmt.Println(g == nil)

	c = g
	fmt.Println(c == nil)
	fmt.Printf("c: %T, %v\n", c, c)
}

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (p *Gopher) useLanguage(language string) {
	p.language = language
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}
