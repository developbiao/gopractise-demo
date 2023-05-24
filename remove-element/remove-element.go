package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}
    index := 3
    // remove the element at index i from a.
    copy(slice[index:], slice[index+1:]) 
    fmt.Println("Copy data:", slice)
    slice = slice[:len(slice) - 1]
    fmt.Println("Result:", slice)
    fmt.Println("Main func done!")
}
