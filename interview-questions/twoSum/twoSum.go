package main

import "fmt"

func twoSum(nums []int, target int) []int {
    length := len(nums)
    for i := 0; i < length; i++ {
        for j := i + 1; j < length; j++ {
            if target == nums[i] + nums[j] {
                return []int{i, j}
            }
        }
    }
    return nil
}

func main() {
    result := twoSum([]int{2, 7, 11, 15}, 27)
    fmt.Println(result)
    fmt.Println("Main func done!")
}
