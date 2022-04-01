package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("Time 2 stopped")
	}

	time.Sleep(6 * time.Second)

	fmt.Println("Programing is ok!")
}
