package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// Delete element 7 in slice
	slice = append(slice[:6], slice[7:]...)
	fmt.Println(slice)
}
