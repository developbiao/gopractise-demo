package main

import "fmt"

type Set map[string]struct{}

func main() {
	set := make(Set)
	for _, item := range []string{"Apple", "Pear", "Banana"} {
		set[item] = struct{}{}
	}
	fmt.Println(len(set))

	if _, ok := set["Apple"]; ok {
		fmt.Println("Apple exists")
	}
}
