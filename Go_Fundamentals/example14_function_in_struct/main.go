package main

import "fmt"

type Account struct {
	username, password string
}

type product struct {
	name  string
	price int
	stock int
}

type rectangle struct {
	w, l float64
}

func main() {
	a := Account{"admin", "1234"}
	fmt.Println(a)
	a.clear()
	fmt.Println(a)
	a.print()

	//-------------------------------------

	var p1 product
	p1.name = "Arduino"
	p1.price = 100
	p1.stock = 20

	show(p1)
	// p1 = p1.clear()
	p1 = p1.setDiscount(1)
	show(p1)

	//-------------------------------------

	rec := rectangle{
		w: 4,
		l: 5,
	}

	fmt.Println(rec.l)
	fmt.Println(rec.w)

	rec.SetWidth(6)

	fmt.Println(Area(rec))
	fmt.Println(rec.Area())

}

func (acc *Account) clear() {
	acc.username = ""
	acc.password = ""
}

func (acc Account) print() {
	fmt.Println(acc)
}

//-------------------------------------

func show(p product) {
	fmt.Println(p)
}

func (p product) setDiscount(d int) product {
	p.price = p.price - d
	return p
}

// เพิ่ม func เข้าไปใน struct
// func (p product) clear() product {
// 	p.stock = 0
// 	p.price = 0
// 	return p
// }

//-------------------------------------

func Area(rec rectangle) float64 {
	return rec.l * rec.w
}

func (rec rectangle) Area() float64 {
	return rec.l * rec.w
}

func (rec *rectangle) SetWidth(w float64) {
	rec.w = w
}
