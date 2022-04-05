package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{3, 1, 8}
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("ints:	", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)

	// Reverse functino returns an slice in the reverse order
	float64s := []float64{2.2, 1.2, 3.3, 10, 9.2, 1.3}
	sort.Sort(sort.Reverse(sort.Float64Slice(float64s)))
	fmt.Println("Floats:", float64s)
}
