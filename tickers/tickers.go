package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Receive done mssage!")
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}

	}()

	time.Sleep(3 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

}
