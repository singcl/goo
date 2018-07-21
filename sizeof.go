package main

import (
	"fmt"
	"unsafe"
)

type Man struct {
	Name int
	Age  int
}

func main() {

	m := Man{Name: 2, Age: 20}

	fmt.Println("man size:", unsafe.Sizeof(m))
	fmt.Println("name size:", unsafe.Sizeof(m.Name))
	fmt.Println("age size:", unsafe.Sizeof(m.Age))
}
