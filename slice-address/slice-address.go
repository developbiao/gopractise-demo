package main

import "fmt"

func main() {
    myWeight := []int{11, 12, 13, 14, 15, 16, 17}
    fmt.Printf("myWeight: %v\n", myWeight)
    fmt.Printf("Address of myWeight: %p %p\n", myWeight, &myWeight)
    fmt.Printf("myWeight len: %v, cap: %v\n\n", len(myWeight), cap(myWeight))

    // Reset weight
    resetWeight(myWeight)

    // Add data
    addWeightRecord(myWeight)
    fmt.Printf("myWeight: %v\n", myWeight)
    fmt.Printf("Address of myWeight: %p %p\n", myWeight, &myWeight)
    fmt.Printf("myWeight len: %v, cap: %v\n\n", len(myWeight), cap(myWeight))

}

func resetWeight(weight []int) {
    for i := 0; i < len(weight); i++ {
        weight[i] = weight[i] + 1 * 10
    }
    fmt.Printf("Address of weight: %p   %p\n\n", weight, &weight)
}

func addWeightRecord(weight []int) {
    weightCap := cap(weight)
    weight[0] = 10
    fmt.Printf("Cap of weight: %v\n\n", weightCap)
    for i := 0; i < weightCap-1; i++ {
        weight = append(weight, i)

    }
    fmt.Printf("Weight: %v\n", weight)
    fmt.Printf("Address of Weight: %p   %p\n", weight, &weight)
    fmt.Printf("Weight: %v, cap: %v\n\n", len(weight), cap(weight))

    // Extended slice 
    for i := 0; i < 3; i ++ {
        weight = append(weight, i)
    }
    fmt.Printf("Extended weight: %v\n", weight)
    fmt.Printf("Address of extend weight: %p    %p\n", weight, &weight)
    fmt.Printf("Extended weight len: %v, cap: %v\n\n", len(weight), cap(weight))
}



