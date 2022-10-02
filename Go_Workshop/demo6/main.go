package main

import "fmt"

func main() {
	num := 2
	double(&num)
	fmt.Println(num)

	name := "Bob"
	appendGreeting(&name)
	fmt.Println(name)
}

// n := 2
// double(&n)
// n == 4
func double(n *int) {
	*n = *n * 2
}

// name := "Bob"
// appendGreeting(&name)
// name == "Hi, Bob"
func appendGreeting(s *string) {
	*s = fmt.Sprintf("Hi, %s", *s)
}
