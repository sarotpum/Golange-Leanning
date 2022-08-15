package main

import "fmt"

type rectangle struct {
	w, l float64
}

type product struct {
	name  string
	price int
	stock int
}

func main() {
	rec := rectangle{
		w: 4,
		l: 5,
	}

	fmt.Println(rec.l)
	fmt.Println(rec.w)
	fmt.Println(Area(rec))

	// ---------------------------

	var p1 product
	p1.name = "Arduino"
	p1.price = 100
	p1.stock = 20
	show(p1)
	update(&p1)
	show(p1)
}

func Area(rec rectangle) float64 {
	return rec.l * rec.w
}

func show(p product) {
	fmt.Println(p)
}

func update(p *product) {
	// fmt.Println(p)
	// fmt.Println(*p)
	p.price = p.price + 100
	p.stock = 10
}
