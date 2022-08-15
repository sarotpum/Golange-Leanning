package main

import "fmt"

func main() {
	msg := "some message"
	var msgPointer *string = &msg //* => ใส่เพื่อว่าจะใส่ค่าใน pointer, & => pointer ตัวแปรที่เราจะใส่
	fmt.Println(msg)
	fmt.Println(msgPointer)
	fmt.Println(*msgPointer) // * เพื่อจะนำค่า value ออกมาแสดง ex. = some message

	changeMessage(&msg, "new message 1")
	fmt.Println(msg)

	changeMessage(msgPointer, "new message 2")
	fmt.Println(msg)

	changeMessage(msgPointer, "new message 3")
	fmt.Println(*msgPointer)
}

func changeMessage(aPointer *string, newMessage string) {
	*aPointer = newMessage
}

/* demo2

func main() {

	i := 42
	p := &i
	var p2 *int
	p2 = p

	fmt.Println("I p = ", *p)
	fmt.Println("I p2 = ", *p2)
	fmt.Println("Address of I = ", p)
	*p = *p / 2
	fmt.Println("I = ", *p)

	mutate(p)
	fmt.Println("I = ", *p)
}

func mutate(_p *int) {
	*_p = 25
}

*/


/* demo3

func main() {
	d := 2
	double(&d)
	fmt.Println(d)
}

func double(d *int) {
	*d = *d * 2
}

*/