package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func main() {
	c := new(Customer)
	c.Name = "Singcl"
	c.log = new(Log)
	c.log.msg = "1- Yes I am Come!"
	// shorter
	c = &Customer{"Barach Obama", &Log{"1-Yes wE COME"}}
	fmt.Println(c) //&{Barak Obama 1 - Yes we can!}
}
