package main

import "fmt"

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d\n", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {

	e := Employee{
		firstName:   "Jam",
		lastName:    "Ramanathan",
		basicPay:    13000,
		pf:          2000,
		totalLeaves: 30,
		leavesTaken: 5,
	}

	var s SalaryCalculator = e
	s.DisplaySalary()

	var l LeaveCalculator = e
	fmt.Println("Leaves left =", l.CalculateLeavesLeft())
}
