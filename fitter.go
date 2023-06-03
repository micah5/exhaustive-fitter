package main

import (
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"math"
)

func PlotShapes(shape1, shape2 []Point, path string) {
	p := plot.New()

	addShapeToPlot(p, shape1, plotter.Polygon{Color: color.RGBA{R: 255, A: 255}})
	addShapeToPlot(p, shape2, plotter.Polygon{Color: color.RGBA{G: 255, A: 255}})

	if err := p.Save(10*vg.Inch, 10*vg.Inch, path); err != nil {
		panic(err)
	}
}

func addShapeToPlot(p *plot.Plot, shape []Point, lineStyle plotter.Polygon) {
	pts := make(plotter.XYs, len(shape))
	for i := range shape {
		pts[i].X = shape[i].X
		pts[i].Y = shape[i].Y
	}

	polygon, err := plotter.NewPolygon(pts)
	if err != nil {
		panic(err)
	}

	polygon.Color = lineStyle.Color
	p.Add(polygon)
}

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

func Circle(w, h float64, resolution int) []Point {
	vertices := make([]Point, 0)
	for i := 0; i < resolution; i++ {
		angle := float64(i) / float64(resolution) * math.Pi * 2
		vertex := Point{math.Cos(angle)*w/2.0 + w/2.0, math.Sin(angle)*h/2.0 + h/2.0}
		vertices = append(vertices, vertex)
	}
	return vertices
}

func main() {
	// Example usage.
	src := [4]Point{
		{0, 0},
		{1, 0},
		{1, 1},
		{0, 1},
	}
	dst := [4]Point{
		{0, 2},
		{1, 1},
		{1, 0},
		{0, 1},
	}
	transform := computeProjectiveTransform(src, dst)
	//fmt.Println(mat.Formatted(&transform))

	// Transform each source point and print the result.
	shape := Circle(1, 1, 8)
	transformed := make([]Point, 0)
	for _, pt := range shape {
		transformedPt := transformPoint(pt, &transform)
		//fmt.Println("Original point:", pt)
		//fmt.Println("Transformed point:", transformedPt)
		transformed = append(transformed, transformedPt)
	}

	PlotShapes(dst[:], transformed, "test.png")
}
