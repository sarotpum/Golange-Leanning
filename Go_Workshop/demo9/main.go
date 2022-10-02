package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Apple Banana Apple Banana apple"
	fmt.Println(wordCount(input))
}

// TODO: split words and count them
func wordCount(s string) map[string]int {
	split := strings.Split(s, " ")
	result := map[string]int{}
	for i := 0; i < len(split); i++ {
		result[split[i]] = result[split[i]] + 1
	}

	return result
}
