package main

import "fmt"

func main() {
	number := make(chan bool)
	letter := make(chan bool)
	done := make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()

	go func() {
		j := 'A'
		for {
			select {
			case <-letter:
				fmt.Print(string(j))
				j++
				fmt.Print(string(j))
				j++
				number <- true

				if j >= 'Z' {
					done <- true
				}
			}
		}
	}()

	number <- true

	<-done
	fmt.Println()
	fmt.Println("Main func done!")
}
