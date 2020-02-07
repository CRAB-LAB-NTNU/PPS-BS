package plotter

import (
	"image/color"
	"log"
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

//Plot Takes in two lists of individuals. The first list is the population while the second is the archive
func (plotter2d *Plotter2D) Plot(popululation, archive []types.Individual) {

	points, maxX, maxY := convertIndividualsToPoints2D(popululation)
	points2, maxX2, maxY2 := convertIndividualsToPoints2D(archive)

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}

	p.Title.Text = plotter2d.Title
	p.X.Label.Text = plotter2d.LabelX
	p.X.Label.Text = plotter2d.LabelY
	p.X.Max = math.Max(maxX, maxX2) + 0.2
	p.Y.Max = math.Max(maxY, maxY2) + 0.2

	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(points)
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(1)

	a, err := plotter.NewScatter(points2)
	a.GlyphStyle.Color = color.RGBA{R: 0, B: 255, A: 255}
	a.GlyphStyle.Radius = vg.Points(1)

	p.Add(s)
	p.Add(a)
	p.Legend.Add("Archive", a)
	p.Legend.Add("Population", s)
	err = p.Save(500, 500, "plotter/testdata/scatter.png")
	if err != nil {
		log.Panic(err)
	}

}

func (plotter2d *Plotter2D) PlotInfeasible(pop []types.Individual) {
	var points plotter.XYs
	for i := range pop {
		if !feasible(pop[i].Fitness()) {
			ov := pop[i].Fitness().ObjectiveValues
			point := plotter.XY{X: ov[0], Y: ov[1]}
			points = append(points, point)
		}
	}
	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = plotter2d.Title
	p.X.Label.Text = plotter2d.LabelX
	p.Y.Label.Text = plotter2d.LabelY
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(points)
	s.GlyphStyle.Color = color.RGBA{R: 0, G: 0, B: 0, A: 100}
	s.GlyphStyle.Radius = vg.Points(1)
	p.Add(s)
	p.Save(200, 200, "plotter/testdata/region.png")
}

func violation(f types.Fitness) float64 {
	var s float64
	for _, f := range f.ConstraintValues {
		if f < 0 {
			s += math.Abs(f)
		}
	}
	return s
}

func feasible(f types.Fitness) bool {
	return violation(f) <= 0
}

func convertIndividualsToPoints2D(individuals []types.Individual) (plotter.XYs, float64, float64) {
	points := make(plotter.XYs, len(individuals))
	maxX := math.SmallestNonzeroFloat64
	maxY := math.SmallestNonzeroFloat64
	for i := range points {
		objectiveValues := individuals[i].Fitness().ObjectiveValues
		points[i].X = objectiveValues[0]
		points[i].Y = objectiveValues[1]
		if objectiveValues[0] > maxX {
			maxX = objectiveValues[0]
		}
		if objectiveValues[1] > maxY {
			maxY = objectiveValues[1]
		}
	}
	return points, maxX, maxY
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
