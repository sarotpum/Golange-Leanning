package main

import "github.com/fatih/color"

func main() {
	// fmt.Printf("Hello module 1\n")
	color.Blue("%s \n", "Hello module 1")
	color.Red("We have red")
	color.Yellow("Yellow color too!")
	color.Magenta("And many others ..")
}
