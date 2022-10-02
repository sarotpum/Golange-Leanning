package main

import "fmt"

func main() {

	fmt.Printf("sumOfFirst => %d\n", sumOfFirst(3))

	result := isPrime(4)

	fmt.Println(result)
}

// It should return sum of n, n-1, n-2, ..., 1
// sumOfFirst(3) should return 3+2+1=6
func sumOfFirst(n int) int {
	sum := 0
	for i := n; i > 0; i-- {
		sum = sum + i
	}

	return sum
}

// A prime number is a number greater than 1 with only two factors – themselves and 1
func isPrime(n int) bool {
	for i := 2; i < n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return n > 1 // 1 is not a prime number
}
