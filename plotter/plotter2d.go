package plotter

import (
	"bufio"
	"image/color"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

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
	p.X.Label.Text = plotter2d.LabelY

	p.X.Max = plotter2d.Max
	p.X.Min = plotter2d.Min
	p.Y.Max = plotter2d.Max
	p.Y.Min = plotter2d.Min
	p.Add(plotter.NewGrid())

	if plotter2d.Solution != nil {
		pareto := convertParetoToPoints(plotter2d.Solution, plotter2d.Max)
		paretoScatter, _ := plotter.NewScatter(pareto)

		paretoScatter.GlyphStyle.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
		paretoScatter.GlyphStyle.Radius = vg.Points(0.5)

		p.Add(paretoScatter)
		p.Legend.Add("Solution", paretoScatter)
	}

	if plotter2d.Extremes != nil {
		extr := convertParetoToPoints(plotter2d.Extremes, plotter2d.Max)

		extrScatter, _ := plotter.NewScatter(extr)
		extrScatter.GlyphStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
		extrScatter.GlyphStyle.Radius = vg.Points(2)

		p.Add(extrScatter)
	}

	for i, pop := range pops {
		x := convertIndividualsToPoints2D(pop, plotter2d.Max)
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

func convertParetoToPoints(pareto [][]float64, max float64) plotter.XYs {
	var points plotter.XYs
	for _, set := range pareto {
		if set[0] > max || set[1] > max {
			continue
		}
		point := plotter.XY{
			X: set[0],
			Y: set[1],
		}
		points = append(points, point)
	}
	return points
}

func convertIndividualsToPoints2D(individuals []types.Individual, max float64) plotter.XYs {
	var points plotter.XYs
	for _, ind := range individuals {
		objectiveValues := ind.Fitness().ObjectiveValues
		if objectiveValues[0] > max || objectiveValues[1] > max {
			continue
		}
		point := plotter.XY{
			X: objectiveValues[0],
			Y: objectiveValues[1],
		}
		points = append(points, point)
	}
	return points
}

func ParseDatFile(path string) [][]float64 {
	var bucket [][]float64
	if file, err := os.Open(path); err != nil {
		log.Fatal(err)
	} else {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			cords := strings.Fields(line)
			var point []float64
			for _, f := range cords {
				if v, err := ParseFloat(f); err != nil {
					log.Fatal(err)
				} else {
					point = append(point, v)
				}
			}
			bucket = append(bucket, point)
		}
	}
	return bucket
}

func ParseFloat(str string) (float64, error) {
	val, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return val, nil
	}

	//Some number may be seperated by comma, for example, 23,120,123, so remove the comma firstly
	str = strings.Replace(str, ",", "", -1)

	//Some number is specifed in scientific notation
	pos := strings.IndexAny(str, "eE")
	if pos < 0 {
		return strconv.ParseFloat(str, 64)
	}

	var baseVal float64
	var expVal int64

	baseStr := str[0:pos]
	baseVal, err = strconv.ParseFloat(baseStr, 64)
	if err != nil {
		return 0, err
	}

	expStr := str[(pos + 1):]
	expVal, err = strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return baseVal * math.Pow10(int(expVal)), nil
}
