package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4 ,5}
    reverse(slice)
    fmt.Println(slice)
    var a [3]int
    a = [...]int{1, 2, 3}
    fmt.Println(a)
    fmt.Printf("%v\n", reverse([]int{6, 7, 8, 9}))
    fmt.Println("Main func done!")
}


func reverse(s []int) []int {
    for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
        s[i], s[j] = s[j], s[i]
    }
    return s
}

