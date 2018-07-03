package main

import (
	"fmt"

	"github.com/singcl/go-simple/pack1"
)

func main() {
	var test1 string
	test1 = pack1.ReturnStr()
	test2 := pack1.ReturnStr2()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("ReturnStr from package1: %s\n", test2)
	fmt.Printf("Integer from package1: %d\n", pack1.Pack1Int)
	// fmt.Printf("Float from package1: %f\n", pack1.pack1Float)
}
