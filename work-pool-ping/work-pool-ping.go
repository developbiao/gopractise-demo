package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"
)

type Job struct {
	id   int
	host string
}

type Result struct {
	job   Job
	pong  string
	alive bool
}

// job channel
var jobChannel = make(chan Job, 10)

// result channel
var resultChannel = make(chan Result, 10)

// alive hosts
var aliveHosts []string

func createWorkerPool(numberOfWorker int) {
	var wg sync.WaitGroup
	for i := 0; i < numberOfWorker; i++ {
		wg.Add(1)
		go work(&wg)
	}
	wg.Wait()
	close(resultChannel)
}

func allocateJob(numberOfJob int) {
	for i := 1; i <= numberOfJob; i++ {
		job := Job{id: i, host: "192.168.61." + fmt.Sprint(i)}
		jobChannel <- job
	}
	close(jobChannel)
	fmt.Println("Job already allocate to channel")
}

// Worker
func work(wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobChannel {
		alive, pong := pingServer(job.host)
		output := Result{job: job, alive: alive, pong: pong}
		resultChannel <- output
	}
}

func pingServer(host string) (alive bool, pong string) {
	cmd := exec.Command("ping", "-c", "1", host)
	err := cmd.Run()
	var check string
	if err != nil {
		check = fmt.Sprintf("Error pinging %s: %s\n", host, err)
		alive = false
	} else {
		check = fmt.Sprintf("%s is alive \n", host)
		alive = true
	}
	return alive, check
}

func getResult(done chan<- bool) {
	for result := range resultChannel {
		fmt.Printf("Job id %d, ping %s result: %s\n", result.job.id, result.job.host, result.pong)
		if result.alive {
			aliveHosts = append(aliveHosts, result.job.host)
		}
	}
	done <- true
}

func main() {
	startTime := time.Now()
	const numberOfJob int = 254
	const numberOfWorker int = 200

	// Allocate job
	go allocateJob(numberOfJob)

	// Get ping server result
	done := make(chan bool)
	go getResult(done)

	// Create worker pool
	createWorkerPool(numberOfWorker)
	<-done

	fmt.Println("---- Alive host list -----")
	for _, aliveHost := range aliveHosts {
		fmt.Println(aliveHost)
	}

	diff := time.Now().Sub(startTime)
	fmt.Println("Total time taken ", diff.Seconds(), "seconds for ping")
}
