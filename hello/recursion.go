package main

import "fmt"

func main () {

    result := fibonacci(20)
    fmt.Println("result", result)
}

func fibonacci(n int) int {
    if n == 0 {
        return 1
    }
    return n * fibonacci(n - 1)
}
