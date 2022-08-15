package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// example.1
	example1()

	// example.2
	example2()
}

func example2() {
	wg := &sync.WaitGroup{}

	wg.Add(3)

	start := time.Now()
	go doSomething(wg)
	go doSomething(wg)
	go doSomething(wg)
	// time.Sleep(120 * time.Millisecond)

	wg.Wait()

	fmt.Println(time.Since(start))

	// example()
}

func example1() {
	go run1()
	go run2()

	time.Sleep(5 * time.Second)
}

func run1() {
	for i := 0; i < 100; i++ {
		fmt.Println("Run1 something")
	}
}

func run2() {
	for i := 0; i < 100; i++ {
		fmt.Println("Run2 something")
	}
}

func doSomething(wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("something")
}

func example() {
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Print("-")
		}
	}()
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			fmt.Print("+")
		}
	}()

	time.Sleep(time.Second)
}