package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address2 struct {
	Type    string
	City    string
	Country string
}

type VCard2 struct {
	FirstName string
	LastName  string
	Addresses []*Address2
	Remark    string
}

var content2 string
var vc VCard2

// 读取gob 文件
func main() {
	// using a decoder:
	file, _ := os.Open("vcard.gob")
	defer file.Close()

	// dec := gob.NewDecoder(file) // 不适用bufio 这样也可以创建dec
	// 为什么使用bufio 请看bufio 的相关接口说明

	inReader := bufio.NewReader(file)
	dec := gob.NewDecoder(inReader)

	err := dec.Decode(&vc)

	if err != nil {
		log.Println("Error in decoding gob")
	}
	fmt.Println(vc)
}
