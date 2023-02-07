package main

import "fmt"

type SalaryCalculator interface {
    CalculateSalary() int
}

type Permanent struct {
    empId       int
    basicpay    int
    pf          int
}

type Contract struct {
    empId       int
    basicpay    int
}

type Freelancer struct {
    empId       int
    ratePerHour int
    totalHours  int
}

// Salary of permanent employee is the sum of baisc pay and pf
func (p Permanent) CalculateSalary() int {
    return p.basicpay + p.pf
}

// Salary of Contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
    return c.basicpay
}

// Salary of freelancer
func (f Freelancer) CalculateSalary() int {
    return f.ratePerHour * f.totalHours
}

/*
total expense is calculated by iterating through the SalaryCalcuator
slice and summing
the slaries of the individual employees
*/
func totalExpense(s []SalaryCalculator) int {
    expense := 0
    for _, v := range s {
        expense += v.CalculateSalary()
    }
    return expense    
}

func main() {
    pemp1 := Permanent{
        empId:    1,
        basicpay: 5000,
        pf:       20,
    }
    pemp2 := Permanent{
        empId:    2,
        basicpay: 6000,
        pf:       30,
    }
    cemp1 := Contract{
        empId:    3,
        basicpay: 3000,
    }
    freelancer1 := Freelancer {
        empId: 4,
        ratePerHour: 70,
        totalHours: 120,
    }
    freelancer2 := Freelancer {
        empId: 5,
        ratePerHour: 100,
        totalHours: 100,
    }


    employees := []SalaryCalculator{pemp1, pemp2, cemp1, freelancer1, freelancer2}
    totalExpense := totalExpense(employees)
    fmt.Printf("Total Expense Per Month $%d\n", totalExpense)
}
