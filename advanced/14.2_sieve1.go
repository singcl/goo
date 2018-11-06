// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.package main
package main

import "fmt"

//Generate产生2-N的序列依次写入chan，main中循环第一次读到2，然后调用Filter(ch, ch1, 2)，这个Filter读取3-N，并从中删除2的倍数的数，将它输出到ch1，
//注意ch = ch1这一行，此时main中的ch的输入是第一轮循环中的Filter的输出，即过滤掉2的倍数的数的序列。
//第二轮循环开始时，ch序列是3,5,7,9,11,13,15,.....，先输出3，然后删除这个序列中3的倍数，将输出串联到下一次循环的输入。
//一个删除了所有比它小的素数的倍数的数的序列，就是一个素数序列。

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in, out chan int, prime int) {
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func main() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}