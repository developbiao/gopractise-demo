package main

import "fmt"

// IGreeting duck type
type IGreeting interface {
	sayHello()
}

func sayHello(i IGreeting) {
	i.sayHello()
}

type Golang struct {
}

func (g Golang) sayHello() {
	fmt.Println("Hi, I am Golang!")
}

type PHP struct {
}

func (p PHP) sayHello() {
	fmt.Println("Hi, I am PHP!")
}

func main() {
	golang := Golang{}
	php := new(PHP)

	sayHello(golang)
	sayHello(php)

}
