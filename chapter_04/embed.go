package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}
type Circle struct {
	Center Point
	Radius int
}
type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {

	var w Wheel

	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Center: Point{X: 8, Y: 8},
			Radius: 5},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}

	fmt.Printf("%#v\n", w)
	w.Circle.Center.X = 42
	fmt.Printf("%#v\n", w)

}
