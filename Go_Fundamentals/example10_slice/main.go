package main

import "fmt"

func main() {
	prims := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = prims[3:6]
	fmt.Println(s)

	slicePrimes := []int{2, 3, 5, 7, 11, 13, 17, 19}

	printSlice(slicePrimes)
	printSlice(slicePrimes[2:]) // leading slice
	printSlice(slicePrimes[:5]) // Trailing slice
	printSlice(slicePrimes[3:5])

	slicePrimesPointer := &slicePrimes
	fmt.Println(slicePrimesPointer)
	slicePrimes = append(slicePrimes, 23)
	printSlice(slicePrimes)

	// slicePrimes = append(slicePrimes, 23, 29)
	// printSlice(slicePrimes)

	slicePrimes = append(slicePrimes, []int{23, 29}...)
	printSlice(slicePrimes)

	test2 := "test"
	fmt.Println(test2[0:2])

	//-------------------------------------
	var numbers1 = make([]int, 3, 5) // 3 len, 5 cap
	numbers1 = append(numbers1, 1)
	numbers1 = append(numbers1, 2)
	showSlice(numbers1)

	var numbers2 []int
	numbers2 = append(numbers2, 1)
	numbers2 = append(numbers2, 2)
	numbers2 = append(numbers2, 3)
	showSlice(numbers2)


	//-------------------------------------
	var ss []int

	if ss == nil {
		fmt.Println("it's nil")
	}

	a := [...]int{1, 3, 5, 7, 9} 

	ss = a[:]
	fmt.Printf("%d %d %p %p\n", len(ss), cap(ss), s, &a)
	ss = append(s, 11, 13)
	fmt.Printf("%d %d %p %p\n", len(ss), cap(ss), s, &a)
}

func printSlice(ss []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(ss), cap(ss), ss)
}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
