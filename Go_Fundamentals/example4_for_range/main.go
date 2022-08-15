package main

import "fmt"

func main() {
	data := []string{"a", "b", "c"}
	for i, item := range data {
		fmt.Println(i, item)
	}

	courses := []string{"Android", "ios", "React"}

	for index, item := range courses {
		fmt.Printf("%d. %s\n", index+1, item)
	}

	for _, item := range courses {
		fmt.Printf("%s\n", item)
	}
}