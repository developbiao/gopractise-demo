package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digit(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	// simlate the fact thant it takes some time
	time.Sleep(2 * time.Second)
	return sum
}

func work(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digit(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go work(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{id: i, randomno: randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)

	done := make(chan bool)
	go result(done)

	noOfWorkers := 20
	createWorkerPool(noOfWorkers)
	<-done

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("Total time taken ", diff.Seconds(), "seconds")
}
