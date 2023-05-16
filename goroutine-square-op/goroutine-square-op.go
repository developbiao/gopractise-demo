package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go square(ch2, ch1)
	output(ch2)
	fmt.Println("Main func done!")
}

func counter(in chan<- int) {
	defer close(in)
	for i := 0; i < 100; i++ {
		in <- i
	}
}

func square(in chan<- int, out <-chan int) {
	defer close(in)
	for i := range out {
		in <- i * i
	}
}

func output(out <-chan int) {
	for i := range out {
		fmt.Println(i)
	}
}
