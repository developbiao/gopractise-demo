package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	ch := make(chan int)
	go dummy(ch)
	log.Println("waiting for reception")
	ch <- 45
	fmt.Println("Main func done!")
}

func dummy(c chan int) {
	time.Sleep(3 * time.Second)
	val := <-c
	fmt.Println("Get Value:", val)
}
