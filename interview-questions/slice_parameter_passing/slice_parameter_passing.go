package main

import "fmt"

func main() {
	slice := []int{1, 1, 1}
	newSlice := myAppend(slice)

	fmt.Println(slice)
	fmt.Println(newSlice)

	slice = newSlice
	myAppendPtr(&slice)
	fmt.Println(slice)

	fmt.Println("Main func done!")
}

func myAppend(s []int) []int {
	// Change s but not effect outlet function s
	s = append(s, 100, 200, 300)

	// for i := range s {
	// 	s[i] = 3
	// }
	return s
}

func myAppendPtr(s *[]int) {
	// Change outside s self
	*s = append(*s, 100)
	return
}
