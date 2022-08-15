package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers = map[string]int{"one": 1, "two": 2, "there": 3}
	fmt.Println("", numbers)
	fmt.Println("", numbers["two"])

	var colors = make(map[string]string)
	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	colors["blue"] = "#00f"
	fmt.Println("", colors)
	fmt.Println("", colors["green"])

	var courses = make(map[string]map[string]int)

	courses["Android"] = make(map[string]int)
	courses["Android"]["price"] = 500
	courses["Android"]["price"] = 100
	courses["Android"]["code"] = 1234

	courses["ios"] = make(map[string]int)
	courses["ios"]["price"] = 200
	courses["ios"]["code"] = 444
	fmt.Println(courses)

	fmt.Println(courses["ios"])
	fmt.Println(courses["ios"]["code"])

	//--------------------------------------

	m := map[string]string{}
	//m = make(map[string]string)

	if m == nil {
		fmt.Println("it's nil")
	}

	m["a"] = "apple"
	s := m["a"]
	fmt.Println(s)

	var input = "Apple Banana Apple Banana apple"
	num := wordCount(input)
	fmt.Println(num)
}

func wordCount(s string) map[string]int {
	split := strings.Split(s, " ")
	result := map[string]int{}
	for i := 0; i < len(split); i++ {
		result[split[i]] = result[split[i]] + 1
	}
	return result
}
