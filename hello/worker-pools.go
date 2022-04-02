package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send 9 jobs after close
	for j := 1; j <= 9; j++ {
		fmt.Println("Push job:", j)
		jobs <- j
	}
	close(jobs)

	// Collection results
	for a := 1; a <= 9; a++ {
		<-results
	}

}

// Worker process job from channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}
