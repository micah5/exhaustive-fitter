package main

import (
	"github.com/micah5/exhaustive-fitter"
)

func main() {
	// (The argument to the Square function is padding)
	outer := fitter.Square(0)   // This is the shape we want to fit to
	inner := fitter.Square(0.1) // This is the shape we want to fit (the "inner" shape)
	result, _ := fitter.Transform(inner, outer)
	fitter.Plot("square.png", outer, result) // Helper function to plot the result
}
