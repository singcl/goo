package main

import "fmt"

type Shaper5 interface {
	Area() float32
}

type TopologicalGenus interface {
	Rank() int
}

type Square5 struct {
	side float32
}

func (sq *Square5) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square5) Rank() int {
	return 1
}

type Rectangle5 struct {
	length, width float32
}

func (r Rectangle5) Area() float32 {
	return r.length * r.width
}

func (r Rectangle5) Rank() int {
	return 2
}

func main() {
	r := Rectangle5{5, 3} // Area() of Rectangle needs a value
	q := &Square5{5}      // Area() of Square needs a pointer

	shapes := []Shaper5{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

	topgen := []TopologicalGenus{r, q}
	fmt.Println("Looping through topgen for rank ...")
	for n, _ := range topgen {
		fmt.Println("Shape details: ", topgen[n])
		fmt.Println("Topological Genus of this shape is: ", topgen[n].Rank())
	}
}
