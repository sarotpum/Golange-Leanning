package main

import "fmt"

func main() {
	result := isOdd(3)
	fmt.Println(result)

	// result := false
	// for i := 1; i <= 10; i++ {

	// 	result = isOdd(i)
	// 	fmt.Println(result)
	// }
}

func isOdd(n int) bool {
	if n%2 == 0 {
		return false
	} else {
		return true
	}
}
