package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Microsecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// channel buffering
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// each 200 ms add new value to  burstyLimiter
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// Should be understand how does'it work
	// Mock 5 request
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
