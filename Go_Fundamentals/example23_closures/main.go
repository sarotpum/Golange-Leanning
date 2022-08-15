package main

import "fmt"

func main() {

	fn := newCounterFunc()
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())

	//------------------------

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	seqString := stringSeq()
	fmt.Println(seqString())
	fmt.Println(seqString())

	fmt.Println(stringSeq()())
}

func newCounterFunc() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func stringSeq() func() string {
	y := 0
	return func() string {
		y++
		return fmt.Sprintf("Y = %d", y)
	}
}
