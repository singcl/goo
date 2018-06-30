package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}
	fmt.Printf("YOUR name is %s", input)
	switch input {
	case "Philip\r\n":
		fmt.Println("Welcome Philip!")
	default:
		fmt.Printf("%s You are not welcome here! Goodbye!", input)
	}
}
