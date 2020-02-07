package plotter

import (
	"image/color"
	"log"

	"github.com/CRAB-LAB-NTNU/PPS-BS/pps"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func (plotter2d *Plotter2D) Plot(popululation []pps.Individual) {

	points := convertPopoulationToPoints2D(popululation)

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}

	p.Title.Text = plotter2d.title
	p.X.Label.Text = plotter2d.labelX
	p.X.Label.Text = plotter2d.labelY

	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(points)
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(1)

	p.Add(s)
	p.Legend.Add("Individuals", s)

	err = p.Save(200, 200, "testdata/scatter.png")
	if err != nil {
		log.Panic(err)
	}

}
func convertPopoulationToPoints2D(population []pps.Individual) plotter.XYs {
	points := make(plotter.XYs, len(population))
	for i := range points {
		objectiveValues := population[i].Fitness().ObjectiveValues
		points[i].X = objectiveValues[0]
		points[i].Y = objectiveValues[1]
	}
	return points
}

/*
func main() {

	rnd := rand.New(rand.NewSource(1))

	// randomPoints returns some random x, y points
	// with some interesting kind of trend.
	randomPoints := func(n int) plotter.XYs {
		pts := make(plotter.XYs, n)
		for i := range pts {
			if i == 0 {
				pts[i].X = rnd.Float64()
			} else {
				pts[i].X = pts[i-1].X + rnd.Float64()
			}
			pts[i].Y = pts[i].X + 10*rnd.Float64()
		}
		return pts
	}

	n := 15
	scatterData := randomPoints(n)
	lineData := randomPoints(n)
	linePointsData := randomPoints(n)

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = "Points Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(scatterData)
	if err != nil {
		log.Panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	l, err := plotter.NewLine(lineData)
	if err != nil {
		log.Panic(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	lpLine, lpPoints, err := plotter.NewLinePoints(linePointsData)
	if err != nil {
		log.Panic(err)
	}
	lpLine.Color = color.RGBA{G: 255, A: 255}
	lpPoints.Shape = draw.CircleGlyph{}
	lpPoints.Color = color.RGBA{R: 255, A: 255}

	p.Add(s, l, lpLine, lpPoints)
	p.Legend.Add("scatter", s)
	p.Legend.Add("line", l)
	p.Legend.Add("line points", lpLine, lpPoints)

	err = p.Save(200, 200, "testdata/scatter.png")
	if err != nil {
		log.Panic(err)
	}

}
*/
