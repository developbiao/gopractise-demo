package main

import (
	"fmt"
	"os/exec"
	"sync"
)

var wg sync.WaitGroup

const numJobs = 100
const numWorkers = 50

func worker(id int, jobs <-chan int, results chan<- string) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("worker", id, "starated job", job)
		results <- pingServer(job)
	}
}

func pingServer(serial int) string {
	host := fmt.Sprint(serial)
	cmd := exec.Command("ping", "-c", "1", host)
	err := cmd.Run()
	var check string
	if err != nil {
		check = fmt.Sprintf("Error pinging %s: %s\n", host, err)
	} else {
		check = fmt.Sprintf("%s is alive \n", host)
	}
	return check
}

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results)
	}

	// Assign task
	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	// Get results
	go func() {
		for pong := range results {
			fmt.Println(pong)
		}
	}()

	wg.Wait()

	fmt.Println("Main func is Done!")
}
