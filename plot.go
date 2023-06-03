package fitter

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
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
