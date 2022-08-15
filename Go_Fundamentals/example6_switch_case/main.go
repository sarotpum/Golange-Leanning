package main

import "fmt"

func main() {
	fnSwitchCase()

	fnSwitchCase2()
}

func fnSwitchCase() {
	index := 3
	switch index {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("something else")
	}
}

func fnSwitchCase2() {
	var lim int

	for true {
		fmt.Print("Please enter Limit: ")
		fmt.Scanf("%d", &lim)

		switch lim {
		case 0:
			fmt.Println("Limit is zero")
		case 1:
			fmt.Println("Limit is one")
		case 2:
			fmt.Println("Limit is two")
		default:
			fmt.Println("Limit is somethingelse")
			return
		}
	}

	fmt.Println("Program exit!")
}
