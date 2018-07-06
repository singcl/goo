package main

import (
	"fmt"

	"github.com/singcl/go-simple/person"
)

func main() {
	p := new(person.Person)
	// p.firstName undefined
	// (cannot refer to unexported field or method firstName)
	// p.firstName = "singcl"
	p.SetFirstName("singcl")
	fmt.Println(p.FirstName()) // Output: singcl
}
