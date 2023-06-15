package main

import "fmt"

func main() {
	arr := []int{64, 34, 23, 15, 13}
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
