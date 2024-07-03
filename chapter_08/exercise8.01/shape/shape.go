/***
Separating the polymorphism code into packages now
***/

package shape

import "fmt"

// declaring the Shape interface
type Shape interface {
	area() float64
	name() string
}

// exporting the shapes
type Triangle struct {
	Base   float64
	Height float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Square struct {
	Side float64
}

// implementing the Shape interface for each shape
func (t Triangle) area() float64 {
	return (0.5 * t.Base) * t.Height
}
func (t Triangle) name() string {
	return "Triangle"
}

func (r Rectangle) area() float64 {
	return r.Length * r.Width
}
func (r Rectangle) name() string {
	return "Rectangle"
}

func (s Square) area() float64 {
	return s.Side * s.Side
}
func (s Square) name() string {
	return "Square"
}

// calculate area and print names for multiple shapes using the interface
func PrintShapeDetails(shape ...Shape) {
	for _, item := range shape {
		fmt.Printf("The %s has an area: %.2f\n", item.name(), item.area())
	}
}
