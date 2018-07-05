// 结构体类型上的方法
package main

import (
	"fmt"
)

type TowInts struct {
	a int
	b int
}

func main() {
	tow1 := new(TowInts)
	tow1.a = 12
	tow1.b = 10
	fmt.Printf("The sum is: %d\n", tow1.AddThem())
	fmt.Printf("Add them to the param: %d\n", tow1.AddToParam(20))
}

func (tn *TowInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TowInts) AddToParam(p int) int {
	return tn.a + tn.b + p
}
