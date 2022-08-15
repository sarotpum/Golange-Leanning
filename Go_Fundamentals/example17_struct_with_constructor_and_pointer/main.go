package main

import (
	"github.com/example17/employee"
)

func main() {
	e := employee.Init(
		"Sam",
		"Adolf",
		30,
		20)

	e.LeavesRemaining()

	e = employee.Init(
		"Lek",
		"Adolf",
		30,
		20)

	e.LeavesRemaining()
}
