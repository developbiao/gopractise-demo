package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(time.Second * 2)

	for v := range ch {
		fmt.Println("Read value", v, "from ch")
		time.Sleep(time.Second * 2)
	}

	fmt.Println("Programing runing is kok")

}
