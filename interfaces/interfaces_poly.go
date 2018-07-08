package main

import "fmt"

type Shaper2 interface {
	Area() float32
}

type Square2 struct {
	side float32
}

func (sq *Square2) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square2{5}     // Area() of Square2 needs a pointer

	// shapes := []Shaper2{Shaper(r), Shaper2(q)}
	// or shorter
	shapes := []Shaper2{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
