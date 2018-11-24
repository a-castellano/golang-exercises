// Package geometry defines simple types for plane geometry.
package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"

	p2 := Point{1, 2}
	q2 := Point{4, 6}
	distanceFromP := p2.Distance       // method value
	fmt.Println(distanceFromP(q2))     // "5"
	var origin Point                   // {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.23606797749979", ;5

	scaleP := p2.ScaleBy // method value
	scaleP(2)            // p becomes (2, 4)
	scaleP(3)            //      then (6, 12)
	scaleP(10)           //      then (60, 120)

}
