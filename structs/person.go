package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 1- struct as a value type:
	var pers1 Person
	pers1.firstName = "Charis"
	pers1.lastName = "Singcl"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)
	// 2 - struct as a pointer:
	pers2 := new(Person)
	pers2.firstName = "Charis"
	pers2.lastName = "Singcl"
	(*pers2).lastName = "Singcloooooo"
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)
	// 3 - struct as a literal:
	pers3 := &Person{"Charis", "Singclpppos"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}
