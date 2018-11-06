package main

import (
	"fmt"
	"time"
)

func pump3() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ;i++ {
			ch<- i
		}
	}()
	return ch
}

func suck3(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	stream := pump3()
	go suck3(stream)
	time.Sleep(1e9)
}