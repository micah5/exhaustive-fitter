package main

import (
	"github.com/micah5/exhaustive-fitter"
)

func main() {
	circle := fitter.Circle(100) // argument is resolution
	square := fitter.Square(0)
	result, _ := fitter.Transform(square, circle)
	fitter.Plot("circle.png", circle, result)
}
