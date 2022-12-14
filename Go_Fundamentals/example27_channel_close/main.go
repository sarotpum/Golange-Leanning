package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	go routine1(ch, 10)

	/*
		for true {
			value, ok := <-ch
			if !ok {
				fmt.Println("No more data")
				break
			}

			fmt.Println(value)555
		}
	*/

	for value := range ch {
		fmt.Println(value)
	}
	
	fmt.Println("No more data")
}

func routine1(c chan int, countTo int) {
	for i := 0; i < countTo; i++ {
		c <- i
	}
	close(c)
}
