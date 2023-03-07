package main

import (
	"context"
	"log"
	"runtime"
	"time"
)

func doSth1(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("First goroutine return")
		return
	}
}

func doSth2(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("Second goroutine return")
		return
	}
}

func lanuch() {
	ctx := context.Background()
	// If we do not call cancel, the two last goroutines will be runing indefinitely
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	log.Println("launch first goroutine")
	go doSth1(ctx)

	log.Println("launch second goroutine")
	go doSth2(ctx)
}

func main() {
	log.Println("Begin program")
	go lanuch()
	time.Sleep(time.Millisecond)
	log.Printf("Goroutine count %d", runtime.NumGoroutine())

	// Block
	for {
	}
}
