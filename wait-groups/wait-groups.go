package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		wg.Add(1)
		// If a WaitGropu is explicitly pased into functions, it should be done by pointer
		// Lunch serveral goroutines and increment the WaitGroup counter for each.
		// Avoid re-use of the same i value in each goroutihne closuer.

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

	fmt.Println("Go Programing is ok!")

}
