package main

import "fmt"

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

func main() {
	// Compile failed: Gopher does not implement coder
	// var biao coder = Gopher{"Go"}
	var biao coder = &Gopher{"Go"}
	biao.code()
	biao.debug()
}
