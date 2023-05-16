package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Start goroutine put 0~100 to channel
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// Start goroutine square value
	go func() {
		for {
			i, ok := <-ch1
			if ok {
				ch2 <- i * i
			} else {
				break
			}
		}
		close(ch2)
	}()

	// Main groutine Read value from ch2
	for i := range ch2 {
		fmt.Println(i)
	}

	fmt.Println("Main func done!")
}
