package  main

import (
	"errors"
	"fmt"
)

// PathError records an error and the operation and file path that caused it.
type PathError struct {
	Op string // "open", "unlink", etc.
	Path string // "open", "unlink", etc.
	Err error // Returned by the system call.
}

func (e *PathError) String() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

// 自定义一个错误
var errNotFound = errors.New("not Found error")

func main() {
	fmt.Printf("Error: %v\n", errNotFound)

	//
	if _, err := Sqrt(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

// 自定义开方的一个错误
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math - square root of negative number")
	}
	// implementation of Sqrt
	return 0, nil
}