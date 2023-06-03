package fitter

import (
	"gonum.org/v1/gonum/mat"
	"math"
)

type Point struct {
	X float64
	Y float64
}

// Function to compute the projective transformation matrix given 4 correspondences.
func computeProjectiveTransform(src, dst [4]Point) mat.Dense {
	// Set up the system of equations.
	var equations mat.Dense
	equations.ReuseAs(8, 8)
	for i := 0; i < 4; i++ {
		equations.SetRow(2*i, []float64{src[i].X, src[i].Y, 1, 0, 0, 0, -dst[i].X * src[i].X, -dst[i].X * src[i].Y})
		equations.SetRow(2*i+1, []float64{0, 0, 0, src[i].X, src[i].Y, 1, -dst[i].Y * src[i].X, -dst[i].Y * src[i].Y})
	}

	// Set up the right-hand side.
	var rhs mat.Dense
	rhs.ReuseAs(8, 1)
	for i := 0; i < 4; i++ {
		rhs.SetRow(2*i, []float64{dst[i].X})
		rhs.SetRow(2*i+1, []float64{dst[i].Y})
	}

	// Solve the system of equations.
	var solution mat.Dense
	solution.Solve(&equations, &rhs)

	// Create the 3x3 transformation matrix.
	var transform mat.Dense
	transform.ReuseAs(3, 3)
	transform.SetRow(0, []float64{solution.At(0, 0), solution.At(1, 0), solution.At(2, 0)})
	transform.SetRow(1, []float64{solution.At(3, 0), solution.At(4, 0), solution.At(5, 0)})
	transform.SetRow(2, []float64{solution.At(6, 0), solution.At(7, 0), 1})

	return transform
}

// Function to transform a point using a 3x3 transformation matrix.
func transformPoint(pt Point, transform mat.Matrix) Point {
	// Convert point to homogeneous coordinates.
	homogeneous := mat.NewVecDense(3, []float64{pt.X, pt.Y, 1})

	// Multiply by the transformation matrix.
	var result mat.VecDense
	result.MulVec(transform, homogeneous)

	// Convert back to Cartesian coordinates.
	w := result.At(2, 0)
	return Point{result.At(0, 0) / w, result.At(1, 0) / w}
}

// Compute the area of a quadrilateral given its vertices.
func quadrilateralArea(vertices [4]Point) float64 {
	// Compute the area using the cross product.
	area := 0.0
	for i := 0; i < 4; i++ {
		j := (i + 1) % 4
		area += vertices[i].X*vertices[j].Y - vertices[j].X*vertices[i].Y
	}
	return math.Abs(area) / 2.0
}

// Compute all possible sets of four points and return the one with the largest area.
func largestQuadrilateral(vertices []Point) [4]Point {
	if len(vertices) < 4 {
		panic("Need at least four vertices")
	}
	maxArea := -1.0
	var maxVertices [4]Point
	for i := 0; i < len(vertices); i++ {
		for j := i + 1; j < len(vertices); j++ {
			for k := j + 1; k < len(vertices); k++ {
				for l := k + 1; l < len(vertices); l++ {
					quad := [4]Point{vertices[i], vertices[j], vertices[k], vertices[l]}
					area := quadrilateralArea(quad)
					if area > maxArea {
						maxArea = area
						maxVertices = quad
					}
				}
			}
		}
	}
	return maxVertices
}

func Transform(flatSrc, flatDst []float64) []float64 {
	// Example usage.
	src := [4]Point{
		{0, 0},
		{1, 0},
		{1, 1},
		{0, 1},
	}
	polygon := build(flatDst)
	dst := largestQuadrilateral(polygon)
	transform := computeProjectiveTransform(src, dst)

	// Transform each source point and print the result.
	shape := build(flatSrc)
	transformed := make([]Point, 0)
	for _, pt := range shape {
		transformedPt := transformPoint(pt, &transform)
		transformed = append(transformed, transformedPt)
	}

	return flatten(transformed)
}
