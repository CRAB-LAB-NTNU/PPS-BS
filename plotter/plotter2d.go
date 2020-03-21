package plotter

import (
	"image/color"
	"log"
	"strconv"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

//Plot Takes in two lists of individuals. The first list is the population while the second is the archive
func (plotter2d *Plotter2D) Plot(pops ...[]types.Individual) {
	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = plotter2d.Title
	p.X.Label.Text = plotter2d.LabelX
	p.Y.Label.Text = plotter2d.LabelY
	p.X.Max = plotter2d.Max
	p.X.Min = plotter2d.Min
	p.Y.Max = plotter2d.Max
	p.Y.Min = plotter2d.Min
	p.Add(plotter.NewGrid())
	if plotter2d.Solution != nil {
		pareto := convertParetoToPoints(plotter2d.Solution, plotter2d.Min, plotter2d.Max)
		paretoScatter, _ := plotter.NewScatter(pareto)
		paretoScatter.GlyphStyle.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
		paretoScatter.GlyphStyle.Radius = vg.Points(0.5)
		p.Add(paretoScatter)
		p.Legend.Add("Solution", paretoScatter)
	}
	if plotter2d.Extremes != nil {
		extr := convertParetoToPoints(plotter2d.Extremes, plotter2d.Min, plotter2d.Max)
		extrScatter, _ := plotter.NewScatter(extr)
		extrScatter.GlyphStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
		extrScatter.GlyphStyle.Radius = vg.Points(2)

		p.Add(extrScatter)
	}
	for i, pop := range pops {
		x := convertIndividualsToPoints2D(pop, plotter2d.Min, plotter2d.Max)
		s, _ := plotter.NewScatter(x)
		s.GlyphStyle.Color = color.RGBA{R: 255 * uint8(i), G: 0, B: 0, A: 255}
		s.GlyphStyle.Radius = vg.Points(1)
		p.Add(s)
		p.Legend.Add("Population "+strconv.Itoa(i), s)
	}
	err = p.Save(500, 500, plotter2d.Filename)
	if err != nil {
		log.Panic(err)
	}
}

func (plotter2d *Plotter2D) PlotMetric(data []float64, switchPoint int) {
	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = plotter2d.Title
	p.X.Label.Text = plotter2d.LabelX
	p.Y.Label.Text = plotter2d.LabelY

	p.Y.Max = 0.05
	p.Y.Min = 0

	p.Add(plotter.NewGrid())
	points := convertFloatToPoints(data, switchPoint)
	scatter, _ := plotter.NewLine(points)
	p.Add(scatter)
	err = p.Save(1500, 500, plotter2d.Filename)
	if err != nil {
		log.Panic(err)
	}
}
