package main

import (
	"github.com/micah5/exhaustive-fitter"
)

func main() {
	parallelogram := []float64{
		0, 2,
		1, 1,
		1, 0,
		0, 1,
	}
	circle := fitter.Circle(10)
	result, _ := fitter.Transform(circle, parallelogram)
	fitter.Plot("parallelogram.png", parallelogram, result)
}
