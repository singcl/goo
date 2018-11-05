package main

import "fmt"

//默认情况下，信道是同步且无缓冲的
//通道的发送/接收操作在对方准备好之前是阻塞的
func main() {
    ch1 := make(chan int)
	go pump2(ch1)
    fmt.Println(<-ch1)
}

func pump2(ch chan int) {
	for i := 0;;i++ {
		ch <- i
	}
}