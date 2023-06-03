package main

import (
	"github.com/micah5/exhaustive-fitter"
)

func main() {
	// input is a flat array of x,y coordinates
	outer := []float64{
		0, 0,
		0, 1,
		1, 1,
		1, 0,
	} // This is the shape we want to fit to
	inner := []float64{
		0.25, 0.25,
		0.25, 0.75,
		0.75, 0.75,
		0.75, 0.25,
	} // This is the shape we want to fit (the "inner" shape)
	result, err := fitter.Transform(inner, outer)
	if err != nil {
		// do something
		panic(err)
	}
	fitter.Plot("square.png", outer, result) // Helper function to plot the result
}
