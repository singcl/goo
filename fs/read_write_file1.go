package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputFile := "products.txt"
	outFile := "products_copy.txt"

	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}
