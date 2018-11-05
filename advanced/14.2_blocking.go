package main

import "fmt"

func main() {
	out := make(chan int)
	out <- 3 // 信道的读取默认是阻塞的，所有下面的信道写入协程 不会被执行。所有该信道只读不写：死锁
	go f1(out) // 放到 out <- 3 之前即可正常运行
}

func f1 (ch chan int) {
	fmt.Println(<- ch)
}

// 所有的协程都休眠了 - 死锁！