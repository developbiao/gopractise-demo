package main

import (
	"fmt"
	"sync"
)

type Job struct {
	id int
}

type WorkerPool struct {
	jobs      chan Job
	workersWg *sync.WaitGroup
}

func NewWorkerPool(numWorkers int, maxQueueSize int) *WorkerPool {
	jobs := make(chan Job, maxQueueSize)
	workersWg := &sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		go func() {
			for job := range jobs {
				doWork(job)
			}
			workersWg.Done()
		}()
		workersWg.Add(1)
	}
	return &WorkerPool{jobs: jobs, workersWg: workersWg}
}

func (wp *WorkerPool) AddJob(job Job) {
	wp.jobs <- job
}

func (wp *WorkerPool) Wait() {
	close(wp.jobs)
	wp.workersWg.Wait()
}

func doWork(job Job) {
	fmt.Printf("Starting job %d\n", job.id)
	// Do some work here...
	fmt.Printf("Finished job %d\n", job.id)
}

func main() {
	wp := NewWorkerPool(5, 10)
	for i := 0; i < 20; i++ {
		wp.AddJob(Job{id: i})
	}
	wp.Wait()
}
