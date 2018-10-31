package main

import (
	"os/exec"
	"fmt"
	"os"
)

func main() {
	cmd := exec.Command("cmder")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!", err)
		os.Exit(1)
	}
	fmt.Printf("The command is %v", cmd)
}