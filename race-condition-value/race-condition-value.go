package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	// Solving the race condnition using a mutex
	mutex.Lock()
	x += 1
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, &mutex)
	}
	wg.Wait()
	fmt.Println("Final value of x", x)
}
