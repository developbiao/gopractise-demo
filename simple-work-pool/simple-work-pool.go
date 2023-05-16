package main

import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	workerCount int
	workerChan  chan func()
	wg          sync.WaitGroup
}

func NewWorkerPool(workerCount int) *WorkerPool {
	pool := &WorkerPool{
		workerCount: workerCount,
		workerChan:  make(chan func()),
	}

	for i := 0; i < workerCount; i++ {
		go pool.worker()
	}
	return pool
}

func (pool *WorkerPool) worker() {
	for task := range pool.workerChan {
		task()
	}
	pool.wg.Done()
}

func (pool *WorkerPool) AddTask(task func()) {
	pool.workerChan <- task
}

func main() {
	pool := NewWorkerPool(10)

	for i := 0; i < 100; i++ {
		pool.AddTask(func() {
			fmt.Println("Task", i)
		})
	}

	pool.wg.Wait()

	fmt.Println("Main func done!")
}
