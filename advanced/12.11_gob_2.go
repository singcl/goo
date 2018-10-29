package main

import (
	"encoding/gob"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

var content string

// 信息写进名为 vcard.gob 的文件。
// 这会产生一个文本可读数据和二进制数据的混合，当你试着在文本编辑中打开的时候会看到。
func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"work", "Boom", []*Address{pa, wa}, "none"}

	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// using an encoder:

	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
}
