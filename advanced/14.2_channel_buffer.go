package main

import (
	"time"
	"fmt"
)

// 同步：ch :=make(chan type, value)
// value == 0 -> synchronous, unbuffered (阻塞）
// value > 0 -> asynchronous, buffered（非阻塞）取决于value元素

func main() {
	c := make(chan int, 10) // asynchronous, buffered（非阻塞信道）
	//c := make(chan int, 0) //  synchronous, unbuffered (阻塞信道）
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("Received:", x)
	}()
	fmt.Println("Sending:", 10)
	c <- 10
	fmt.Println("sent:", 10)
}

/* Output:
sending 10
sent 10   // prints immediately
no further output, because main() then stops
*/