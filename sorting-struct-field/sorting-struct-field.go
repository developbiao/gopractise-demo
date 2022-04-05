package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	family := []struct {
		Name string
		Age  int
	}{
		{"Alice", 17},
		{"David", 3},
		{"Eve", 3},
		{"Bob", 28},
	}

	// Using sort by age when age equals keep origional order
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
	fmt.Println(family)

	// Implementing sort by age asc, name desc when age and name equals keep original order
	sort.SliceStable(family, func(i, j int) bool {
		if family[i].Age != family[j].Age {
			return family[i].Age < family[j].Age
		}
		return strings.Compare(family[i].Name, family[j].Name) == 1
	})
	fmt.Println(family)

}
