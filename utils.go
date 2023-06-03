package fitter

import (
	"math"
)

func Circle(resolution int) []float64 {
	vertices := make([]float64, 0)
	w, h := 1.0, 1.0
	for i := 0; i < resolution; i++ {
		angle := float64(i) / float64(resolution) * math.Pi * 2
		x, y := math.Cos(angle)*w/2.0+w/2.0, math.Sin(angle)*h/2.0+h/2.0
		vertices = append(vertices, x, y)
	}
	return vertices
}

func flatten(points []Point) []float64 {
	flatPoints := make([]float64, 0)
	for _, pt := range points {
		flatPoints = append(flatPoints, pt.X, pt.Y)
	}
	return flatPoints
}

func build(flatPoints []float64) []Point {
	points := make([]Point, 0)
	for i := 0; i < len(flatPoints); i += 2 {
		points = append(points, Point{flatPoints[i], flatPoints[i+1]})
	}
	return points
}
