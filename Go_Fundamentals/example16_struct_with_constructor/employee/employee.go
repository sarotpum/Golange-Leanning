package employee

import "fmt"

type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func New(
	firstName string,
	lastName string,
	totalLeaves int,
	leavesTaken int) employee {
	e := employee{
		FirstName:   firstName,
		LastName:    lastName,
		TotalLeaves: totalLeaves,
		LeavesTaken: leavesTaken,
	}
	return e
}

func (e employee) LeavesRemaining() {
	sum := e.TotalLeaves - e.LeavesTaken
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, sum)
}
