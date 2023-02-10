package employee

import "fmt"

type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	return employee{firstName, lastName, totalLeave, leavesTaken}
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
