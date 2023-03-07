package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func doWork(ctx context.Context, resChan chan int) {
	log.Println("[doWork] Launch the doWork")
	sum := 0
	for {
		time.Sleep(time.Millisecond)
		select {
		case <-ctx.Done():
			log.Println("[doWork] ctx Done is recevied inisde doWork")
			return
		default:
			sum++
			if sum > 1000 {
				log.Println("[doWork] sum has reached 1000")
				resChan <- sum
				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[Handler] Request recieved")
		rCtx := r.Context()
		// Create the result channel
		resChan := make(chan int)
		// Launch the function doWrok in a goroutine
		go doWork(rCtx, resChan)
		// Wait for
		// 1. The client drops the conection
		// 2. The function doWrok to finished it works
		select {
		case <-rCtx.Done():
			log.Println("[Handler] context cancled in main handler, client has disconnected")
			return
		case result := <-resChan:
			log.Println("[Handler] Received 1000")
			log.Println("[Handler] Send response")
			fmt.Fprintf(w, "Response %d", result)
		}
	})

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
}
