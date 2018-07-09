package main

import (
	"fmt"
	"math"
)

type Square3 struct {
	side float32
}

func (sq *Square3) Area() float32 {
	return sq.side * sq.side
}

type Circle struct {
	radius float32
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

type Shaper3 interface {
	Area() float32
}

func main() {
	var areaIntf Shaper3
	sq1 := new(Square3)
	sq1.side = 5

	areaIntf = sq1
	// Is Square the type of areaIntf?
	if t, ok := areaIntf.(*Square3); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}

	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}
