package main

import (
	"github.com/example15/employee"
)

func main() {
	e := employee.Employee{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}

	e.LeavesRemaining()
}
