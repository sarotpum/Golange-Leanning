package main

import "fmt"

func main() {
	someValue := 10
	amount := 10

	if someValue == 10 {
		fmt.Println("== 10")
	} else {
		fmt.Println("!= 10")
	}

	if someValue == 10 || someValue < 20 {
		fmt.Println("someValue > 10 || someValue < 2 ")
	} else {
		fmt.Println("Not someValue > 10 || someValue < 2 ")
	}

	if result := doSomething(); result == "OK" {
		fmt.Println("OK")
	} else {
		fmt.Println("NO OK")
	}

	if max := 5; max > amount {
		fmt.Println("max > Amount")
	} else {
		fmt.Println("max =< Amount")
	}
}

func doSomething() string {
	// do something
	return "OK"
}
