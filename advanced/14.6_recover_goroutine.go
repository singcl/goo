package main

import "log"

type Work struct {
	Aage int
}

func main() {
	//
}

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Work failed with %s in %v", err, work)
		}
	}()
	do(work)
}

func do(work *Work) {
	//
}