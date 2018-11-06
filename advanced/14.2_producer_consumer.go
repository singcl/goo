package main

import "fmt"

// integer producer:
func numGen(start, count int, out chan <- int) {
	for i := 0; i < count; i++ {
		out <- start
		start = start + count
	}
	close(out)
}

// integer consumer:
func numEchoRange(in <-chan int, done chan<- bool) {
	// 它从指定通道中读取数据直到通道关闭，才继续执行下边的代码。
	// 很明显，另外一个协程必须写入 ch（不然代码就阻塞在 for 循环了），而且必须在写入完成后才关闭
	for num := range in {
		fmt.Printf("%d\n", num)
	}
	done<- true
}

func main() {
	numChan := make(chan int)
	done := make(chan bool)

	go numGen(0 ,10, numChan)
	go numEchoRange(numChan, done)

	<-done
}