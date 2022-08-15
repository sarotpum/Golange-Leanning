package main

import "fmt"

func main() {
	c := make(chan string)
	quit := make(chan string)
	go func() { // waiting for result from task
		for i := 0; i < 10; i++ {
			fmt.Print(<-c)
		}
		quit <- "end"
	}()

	task(c, quit)
}

func task(c, quit chan string) {
	for {
		select {
		case c <- "Running...":
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
