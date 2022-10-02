package main

import "fmt"

func main() {
	var a interface{}

	a = 10
	fmt.Printf("%T %v\n", a, a)

	a = "ten"
	fmt.Printf("%T %v\n", a, a)

	a = true
	fmt.Printf("%T %v\n", a, a)

	a = func() string { return "hello" }
	fmt.Printf("%T %v\n", a, a)

	fmt.Println()
}
