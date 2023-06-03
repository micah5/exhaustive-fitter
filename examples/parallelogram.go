package main

import (
	"fmt"
	"github.com/micah5/fitter"
)

func main() {
	polygon := fitter.Polygon{Vertices: []fitter.Point{{X: 0, Y: 2}, {X: 1, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: 1}}}

	maxSquare, trans := fitter.MaxFitSquare(polygon)
	fmt.Println("Max square:", maxSquare)
	fmt.Println("Trans:", trans)
}
