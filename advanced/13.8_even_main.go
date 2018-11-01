package main

import "fmt"
import "./even"

func main() {
	for i := 0;  i <= 100; i++ {
		fmt.Printf("iS the Integer %d even? %v\n", i, even.Even(i))
	}
}