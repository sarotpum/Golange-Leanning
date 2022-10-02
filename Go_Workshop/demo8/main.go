package main

import "fmt"

func main() {
	variadicString("a", "b", "c")

	s := []string{"d", "e", "f"}
	variadicString(s...)
}

func variadicString(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
