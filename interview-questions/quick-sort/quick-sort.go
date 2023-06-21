package main

import "fmt"

func quickSort(arr []int, left int, right int) {
	if left < right {
		pivot := partition(arr, left, right)

		// Recursion left and right
		quickSort(arr, left, pivot-1)
		quickSort(arr, pivot+1, right)
	}

}

func partition(arr []int, left int, right int) int {
	// Select right element as pivot
	pivot := arr[right]

	// Define pointer i record less than pivot element position
	i := left - 1

	// Traverse the array, put the elements smaller than the reference value to
	// the left of i, and the elements greater than or equal to the reference value to
	// the right of i
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Put the reference value to the right of i to complete a partition operation
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Sort Before:", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sort After:", arr)
	fmt.Println("Main func done!")
}
