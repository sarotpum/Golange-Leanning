package employee

import "fmt"

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (e Employee) LeavesRemaining() {
	sum := e.TotalLeaves - e.LeavesTaken
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, sum)
}