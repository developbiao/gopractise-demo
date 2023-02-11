package main

import "fmt"

func sum(a, b int) {
	// It's not possible to recofer from panic that has happend ina different goruntine
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from sum func", r)
		}
	}()
	fmt.Printf("a + b = %d\n", a+b)
	done := make(chan bool)
	go devide(a, b, done)
	<-done
}

func devide(a, b int, done chan bool) {
	fmt.Println("Devide result:", a/b)
	done <- true
}

func main() {
	sum(5, 0)
	fmt.Println("Normally returned from main")
}
