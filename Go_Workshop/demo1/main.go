package main

import "fmt"

func main() {
	result1 := greeting("Pallat")
	result2 := greetingWithAge("Pallat", 30)
	result3 := greetingWithAgeAndDrink("Pallat", 30, "Cola")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

func greeting(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func greetingWithAge(name string, age uint) string {
	return fmt.Sprintf("Hello, %s. You are %d years old.", name, age)
}

func greetingWithAgeAndDrink(name string, age uint, drink string) string {
	return fmt.Sprintf("Hello, %s. You are %d years old and your favorite drink is %s.", name, age, drink)
}