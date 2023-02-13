package main

import "fmt"

// Passing functions as arguments to other functions
// func simple(a func(a, b int) int) {
// 	fmt.Println(a(60, 70))
// }

// Returning functions from other functions
func simple() func(a, b int) int {
	f := func(a, b int) int {
		return a * b
	}
	return f
}

func main() {
	// f := func(a, b int) int {
	// 	return a + b
	// }

	// simple(func(a, b int) int {
	// 	return a - b
	// })

	f := simple()
	fmt.Println(f(13, 17))
}
