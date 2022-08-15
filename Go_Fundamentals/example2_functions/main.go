package main

import "fmt"

func main() {
	fmt.Println("1 + 2 = ", add(1, 2))
	fmt.Println(add2(1, 2, "Result: "))

	// ---------------------------------------

	fn1()
	fn2("Good Morning")
	fn2_1("Good Morning", 2)
	fn3("Good Morning", 1)

	fmt.Printf("Sum => %d\n", sum(5, 10))

	const a, b = 2, 3

	//x, y := swap(a, b)
	var x, y int = swap(a, b)
	fmt.Printf("%d, %d => %d, %d\n", a, b, x, y)

	x, y = swap2(10, 20)
	fmt.Printf("%d, %d => %d, %d\n", 10, 20, x, y)

	fmt.Println(split(17))
}

func add(x, y int) int {
	return x + y
}

func add2(x int, y int, msg string) string {
	return fmt.Sprintf("%s %d", msg, x+y)
}

func fn1() {
	fmt.Println("Code Start")
}

func fn2(msg string) {
	fmt.Println(msg)
}

func fn2_1(msg string, version int) {
	fmt.Println(msg, version)
}

func fn3(title string, version int) {
	fmt.Print(title, " ")
	fmt.Println(version)
}

func sum(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}

func swap2(a int, b int) (x, y int) {
	y = a
	x = b
	return
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
