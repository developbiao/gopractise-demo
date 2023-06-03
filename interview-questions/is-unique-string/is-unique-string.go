package main

import (
	"fmt"
	"strings"
)

func main() {
	if isUniqueString01("abccd") {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	fmt.Println("Main func done!")
}

func isUniqueString01(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		// Visible char maximum 128
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}
