package main

import "fmt"

func main() {
	fnFor()
	fnWhile()
	fnWhile2()
	fnWhileUsingBreak()
}

func fnFor() {
	for index := 0; index <= 10; index++ {
		fmt.Printf("For Index %d\n", index)
	}
}

func fnWhile() {
	index := 0
	for index < 5 {
		index++
		fmt.Printf("While-Index %d\n", index)
	}
}

func fnWhile2() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
		if sum > 500 {
			break //return
		}
	}
}

func fnWhileUsingBreak() {
	index := 0
	for true {
		if index > 5 {
			break
		}
		fmt.Printf("WhileBreak-Index %d\n", index)
		index++
	}
}
