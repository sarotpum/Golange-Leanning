package main

import "fmt"

func main() {
	// defer fmt.Println("world")
	// fmt.Println("Hello ")

	for i := 0; i < 10; i++ {
		defer printFinish(i)
	}

	// defer fmt.Println("Finish")

	for i := 0; i < 10; i++ {
		fmt.Println("", i)
	}
}

func printFinish(i int) {
	fmt.Println("Finish: ", i)
}
