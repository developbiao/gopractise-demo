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

type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d\n", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {

	e := Employee{
		firstName:   "Cheng Cheng",
		lastName:    "Zen",
		basicPay:    9000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}

	// Embedding interface
	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("Leaves left = ", empOp.CalculateLeavesLeft())
}
