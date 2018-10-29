package main

import (
	"crypto/md5"
	"fmt"
	"log"
)

func main() {
	hasher := md5.New()
	b := []byte{}

	data := []byte("hello world")
	n, err := hasher.Write(data)

	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}

	fmt.Printf("Result: %x\n", hasher.Sum(b))
}
