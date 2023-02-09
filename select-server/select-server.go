package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go server1(ch1)
	go server2(ch2)
	count := 0
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case s1 := <-ch1:
			fmt.Println(s1)
		case s2 := <-ch2:
			fmt.Println(s2)
		default:
			fmt.Println("No value received")
		}
		count++
		if count >= 8 {
			break
		}
	}
	fmt.Println("main func done!")
}
