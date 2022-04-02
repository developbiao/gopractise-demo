package main

import (
	"fmt"
	"sync"
	"time"
)

// Wait group
var wg sync.WaitGroup

func worker(id int, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worekr", id, "finished job")
		results <- j * 2
	}

}

func main() {
	const numJobs = 10
	const numWorker = 3
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Create 3 workers
	// This starts up 3 workes, initially blocked because thre are no jobs yet
	wg.Add(numWorker)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//  Assign task
	// Here we send 5 jobs and then close that channel to
	// indicate that's all the work we have.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	// Get process results
	// Finally we collect all the results of the work. This also
	// ensures that the worker goroutines have finished. An
	// Aalternative way to wait for multipe goroutines is to use WaitGroup
	// for k := 1; k <= numJobs; k++ {
	// 	oneItem := <-results
	// 	fmt.Println("Got resulst:", oneItem)
	// }

	wg.Wait()
	fmt.Println("All job already finished!")

}
