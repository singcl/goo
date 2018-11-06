package main

import (
	"time"
	"fmt"
)

func pump4() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ;i++ {
			ch<- i
		}
	}()
	return ch
}

func suck4(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

func main() {
	stream := pump4()
	suck4(stream)
	time.Sleep(1e9)
}