package main

import (
	"fmt"
	"time"
)

func helloGo(done chan bool) {
	fmt.Println("Hello Goroutine!")
	time.Sleep(3 * time.Second)
	done <- true
}

func main() {
	done := make(chan bool)
	go helloGo(done)
	<-done
	fmt.Println("Main func received data")
}
