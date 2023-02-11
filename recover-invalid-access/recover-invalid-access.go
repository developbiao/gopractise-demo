package main

import (
	"fmt"
	"runtime/debug"
)

func recoverInvalidSliceAccess() {
	if r := recover(); r != nil {
		fmt.Println("Invalid slice access", r)
		debug.PrintStack()
	}
}

func invalidSliceAccess() {
	defer recoverInvalidSliceAccess()
	slice := []int{1, 8, 9}
	fmt.Println("index of 3 is:", slice[3])
}

func main() {
	invalidSliceAccess()
	fmt.Println("main func done.")
}
