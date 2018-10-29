package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int    "X, Y, Z"
	Name    string "name"
}

type Q struct {
	X, Y *int32
	Name string
}

// 和 JSON 的使用方式一样，Gob 使用通用的 io.Writer 接口，通过 NewEncoder() 函数创建 Encoder 对象并调用 Encode()；
// 相反的过程使用通用的 io.Reader 接口，通过 NewDecoder() 函数创建 Decoder 对象并调用 Decode。
func main() {
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.

	// Encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Singcl"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// Decode (receive) the value.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)

	// "Singcl": {3,4}
}
