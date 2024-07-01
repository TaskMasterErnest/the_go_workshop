/***
Using polymorphism to calculate the area of different shapes
***/

package main

import "fmt"

// create the shape interface
type Shape interface {
	Area() float64
	Name() string
}

// create the shape types; triangle, rectangle, square
type triangle struct {
	base   float64
	height float64
}

type rectangle struct {
	length float64
	width  float64
}

type square struct {
	side float64
}

// create the functions to implement the shape interfaces
func (t triangle) Area() float64 {
	return (0.5 * t.base) * t.height
}
func (t triangle) Name() string {
	return "triangle"
}

func (r rectangle) Area() float64 {
	return r.length * r.width
}
func (r rectangle) Name() string {
	return "rectangle"
}

func (s square) Area() float64 {
	return s.side * s.side
}
func (s square) Name() string {
	return "square"
}

// function that takes the types and passes them to a variadic input interface
func printShapeDetails(shape ...Shape) {
	for _, shape := range shape {
		fmt.Printf("The area of the %s is: %.2f\n", shape.Name(), shape.Area())
	}
}

func main() {
	t := triangle{base: 10.5, height: 12.3}
	r := rectangle{length: 4.4, width: 5.6}
	s := square{side: 7.3}

	printShapeDetails(t, r, s)
}
