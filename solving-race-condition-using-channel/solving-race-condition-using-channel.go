package main

import (
	"fmt"
	"sync"
)

var x int = 0

func increment(ch chan bool, wg *sync.WaitGroup) {
	ch <- true
	x += 1
	<-ch
	wg.Done()
}

func main() {
	ch := make(chan bool, 1)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(ch, &wg)
	}
	wg.Wait()
	fmt.Println("Final value of x", x)
}
