package main

import (
	"fmt"
	"time"
)

var dog = make(chan struct{})
var cat = make(chan struct{})
var fish = make(chan struct{})

func Dog() {
	<-fish
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Dog")
	dog <- struct{}{}
}

func Cat() {
	<-dog
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Cat")
	cat <- struct{}{}
}

func Fish() {
	<-cat
	time.Sleep(100 * time.Millisecond)
	fmt.Println("fish")
	fish <- struct{}{}
}

func main() {
	for i := 0; i < 3; i++ {
		go Dog()
		go Cat()
		go Fish()
	}

	time.Sleep(3 * time.Second)
}
