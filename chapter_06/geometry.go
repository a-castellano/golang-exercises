// Package geometry defines simple types for plane geometry.
package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// A Path is a journey connecting the points with straight lines.
type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0

	for i := 1; i < len(path); i++ {
		sum += path[i-1].Distance(path[i])
	}
	return sum
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call

	fmt.Println(perim.Distance()) // "12"

	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r) // "{2, 4}"

	p2 := Point{1, 2}
	pptr := &p2
	pptr.ScaleBy(2)
	fmt.Println(p2) // "{2, 4}"

	p3 := Point{1, 2}
	(&p3).ScaleBy(2)
	fmt.Println(p3) // "{2, 4}"
	p3.ScaleBy(2)
	fmt.Println(p3) // "{2, 4}"

}
