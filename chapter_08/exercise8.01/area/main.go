package main

// import the shape package
import (
	"github.com/the_go_workshop/chapter_08/exercise8.01/shape"
)

func main() {
	t := shape.Triangle{Base: 15.5, Height: 20.1}
	r := shape.Rectangle{Length: 20, Width: 10}
	s := shape.Square{Side: 10}
	shape.PrintShapeDetails(t, r, s)
}
