package main

import (
	"fmt"

	books "github.com/example32/book"
	"github.com/example32/calculator"
)

func main() {
	fmt.Println(calculator.Add(100, 50))

	b := books.New()
	fmt.Printf("%T %v\n", b, b)

	b.Name = "Rign"
	fmt.Printf("%T %v\n", b, b)
}
