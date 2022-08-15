package main

import (
	"github.com/example16/employee"
)

func main() {
	e := employee.New(
		"Sam",
		"Adolf",
		30,
		20)

	e.LeavesRemaining()
}
