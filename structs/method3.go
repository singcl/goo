package main

import (
	"fmt"
	"math"
)

// type Engine interface {
// 	Start()
// 	Stop()
// }

// type Car struct {
// 	Engine
// }

// func (c *Car) GoToWorkIn() {
// 	// get in car
// 	c.Start()
// 	// drive to work
// 	c.Stop()
// 	// get out of car
// }

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point
	name string
}

func main() {
	n := &NamedPoint{Point{3, 4}, "Singcl"}
	fmt.Println(n.Abs())
	fmt.Println(n.Point.Abs())
}
