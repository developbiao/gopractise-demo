package main 

import "fmt"

// annomayous return paramater create temporary variable
func test() int {
    i := 0
    defer func() {
        fmt.Println("Defer1")
    }()

    defer func() {
        i += 1
        fmt.Println("Defer2")
    }()
    return i
}

func main() {
    fmt.Println(test())
}
