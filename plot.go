package fitter

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"math/rand"
)

func getRandomColor() color.RGBA {
	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func Plot(path string, shapes ...[]float64) {
	p := plot.New()

	for _, shape := range shapes {
		addShapeToPlot(p, build(shape), plotter.Polygon{Color: getRandomColor()})
	}

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
