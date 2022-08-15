package main

import (
	"fmt"
	"math"
	"time"
)

var tmp1 string
var tmp2 bool
var tmp4 = "go programming"
var count int = 0

func main() {
	tmp1 = "codemobiles"
	tmp2 = true
	fmt.Println(tmp1, tmp2)
 
	count++
	fmt.Println("count => ", count)
	count++
	fmt.Println("count => ", count)

	/*
		const tmp4 int = 0 //เปลี่ยนแปลงไม่ได้
		tmp4 = 10
	*/
	tmp5 := 555
	t1, t2, t3 := 1, 2, 3

	fmt.Println("tmp1 = ", tmp1, tmp2, tmp4, tmp5)
	fmt.Println("tx = ", t1, t2, t3)

	const t33 string = "test"
	fmt.Println(t33)

	fmt.Println(math.Exp2(10)) // 1024
	fmt.Println("---------")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished")

}

//cls => clear
//go build => build จะได้ไฟล์ .exe นำไป run ได้เลย
//go-mon index.go
