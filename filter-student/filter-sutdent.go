package main

import "fmt"

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func main() {
	s1 := student{
		firstName: "Jam",
		lastName:  "Sand",
		grade:     "B",
		country:   "USA",
	}

	s2 := student{
		firstName: "Jun",
		lastName:  "Li",
		grade:     "A",
		country:   "China",
	}

	students := []student{s1, s2}
	result := filter(students, func(s student) bool {
		if s.grade == "A" {
			return true
		}
		return false
	})
	fmt.Println("Filter result:", result)

	a := []int{5, 6, 7, 8, 9}
	r := iMap(a, func(n int) int {
		return n * 5
	})
	fmt.Println(r)
}
