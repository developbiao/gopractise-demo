package main

import (
	"fmt"
	"time"
)

func main() {
	//  timer limiter
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This limiter channel will receive a value every 200 milliseconds
	// This is the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("-----------------")

	// We may want to allow short bursts of requests in our rate
	// limiting scheme while preserving the overall rate limit
	// We can accomplish this by buffering our limiter channel
	// This burstyLimiter channel will allow bursts of up to 3 events
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowd bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200 millseconds we'll try to add a new value to bustyLimiter, up to its limit of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Now simulate 5 more incoming request. The first 3 of these
	// Will benefit form the burst capability of burstyLimiter.
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
