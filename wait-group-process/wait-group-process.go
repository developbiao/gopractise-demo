package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	no := 3
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}

	wg.Wait()
	fmt.Println("All go routines finished executing")
}
