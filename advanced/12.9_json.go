package main

import (
	"fmt"
	"encoding/json"
	"os"
	"log"
)

type Address struct {
	Type string
	City string
	Country string
}

type VCard struct {
	FirstName string
	LastName string
	Addresses []*Address
	Remark string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := &VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// go 常用占位符的含义 @see https://studygolang.com/articles/2644
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}

	// JSON format:
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)

	// using an encoder:
	fp, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer fp.Close()

	enc := json.NewEncoder(fp)
	err := enc.Encode(vc)

	if err != nil {
		log.Println("Error in encoding json")
	}
}