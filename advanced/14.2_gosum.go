package main

import "fmt"

func main() {
	c := make(chan int)
	go sum(12, 13, c)
	fmt.Println(<-c)
}

func sum(x, y int, c chan int) {
	c <- x + y
}