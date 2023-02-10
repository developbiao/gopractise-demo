package main

import "oop/employee"

func main() {
	e := employee.New("Jun", "Li", 30, 15)
	e.LeavesRemaining()
}
