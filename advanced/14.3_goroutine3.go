package main

import "fmt"

func main() {
	ch := make(chan string)
	go sendData2(ch)
	getData2(ch)
}

func sendData2(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	close(ch)
}

func getData2(ch chan string) {
	for {
		input, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("%s ", input)
	}
}