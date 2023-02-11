package main

import "fmt"

func divisionByZero() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd:", r)
		}
	}()
	var a int = 7
	var b int = 0
	var c int = a / b
	fmt.Println(c)
}

func main() {
	divisionByZero()
}
