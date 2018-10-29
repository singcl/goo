package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func main() {

	hasher := sha1.New()
	b := []byte{}

	// hasher 第一种方式
	io.WriteString(hasher, "singcl")

	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))

	//
	hasher.Reset()

	// hasher 第二种方式
	data := []byte("We shall overcome!")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	fmt.Printf("Result: %x\n", hasher.Sum(b))
}
