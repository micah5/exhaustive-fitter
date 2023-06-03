package fitter

import (
	"math"
)

func Square(pad float64) []float64 {
	return []float64{
		pad, pad,
		1 - pad, pad,
		1 - pad, 1 - pad,
		pad, 1 - pad,
	}
}

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

func Star() []float64 {
	var points []float64

	// Define the number of points for a 5-sided star
	pointsNum := 5

	// Multiply by 2 because we want to create a star, not a pentagon
	// i.e., we need to also create the inner vertices
	for i := 0; i < pointsNum*2; i++ {
		// Calculate the angle at this step
		// Multiply by 2pi (to convert to radians) and divide by 10 (the total number of points)
		angle := (float64(i) * math.Pi) / 5.0

		// The radius of the circle the points lie on alternates between 0.5 (for the outer vertices)
		// and some inner radius for the inner vertices.
		// Here we use 0.2 for the inner radius to create a reasonably "sharp" star.
		radius := 0.5
		if i%2 == 1 {
			radius = 0.2
		}

		// Calculate the x and y coordinates from the angle
		// Shift by 0.5 to move (0,0) to the center of the star
		x := 0.5 + radius*math.Cos(angle)
		y := 0.5 + radius*math.Sin(angle)

		points = append(points, x, y)
	}

	return points
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
