package pps

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/plotter"
	"github.com/CRAB-LAB-NTNU/PPS-BS/stages"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func (pps *PPS) setIdealAndNadir(generation int) {
	ip, np := biooperators.CalculateNadirAndIdealPoints(pps.MOEA.Population())

	pps.idealPoints[generation] = ip
	pps.nadirPoints[generation] = np
}

func (pps *PPS) changeStage(generation int) bool {
	if pps.currentStage().Type() == types.Push {
		return pps.changePush(generation)
	}
	return pps.currentStage().IsOver()

}

func (pps *PPS) changePush(generation int) bool {
	push, ok := pps.currentStage().(*stages.Push)
	if !ok {
		panic("Could not assert push stage")
	}
	push.Update(generation, pps.MOEA.Population())
	pps.Stages[pps.stage] = push
	return pps.currentStage().IsOver()
}

func (pps *PPS) initStage() {
	if pps.currentStage().Type() == types.Pull {
		pps.initPull()
	}
	//pps.MOEA.CHM().Initialise(pps.MOEA.Generation(), pps.MOEA.MaxViolation())
}

func (pps *PPS) initPull() {
	pps.MOEA.CHM().Initialise(pps.MOEA.Generation(), pps.MOEA.MaxViolation())
}

func (pps PPS) currentStage() types.Stage {
	return pps.Stages[pps.stage]
}

func (pps *PPS) nextStage() {
	pps.stage++
	//If there are no feasible solutions we simply skip the boundary stage
	if pps.currentStage().Type() == types.BinarySearch && len(pps.MOEA.Archive()) == 0 {
		pps.stage++
	}
}

func (pps PPS) plot(generation int) {

	gen := strconv.Itoa(generation)
	eps := strconv.FormatFloat(pps.MOEA.CHM().Threshold(generation), 'E', -1, 64)
	prob := pps.CMOP.Name() + "." + pps.MOEA.CHM().Name()
	stage := pps.currentStage().Name()

	path := "graphics/gif/" + prob

	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatal(err)
	}

	plotter := plotter.Plotter2D{
		Title:    prob + " Stage: " + stage + " gen: " + gen + " eps: " + eps,
		LabelX:   "f1",
		LabelY:   "f2",
		Min:      pps.Export.VideoMin,
		Max:      pps.Export.VideoMax,
		Filename: path + "/" + gen + ".png",
		Solution: pps.paretoPoints,
		Extremes: [][]float64{pps.idealPoints[generation], pps.nadirPoints[generation], pps.MOEA.Ideal()},
	}
	plotter.Plot(pps.MOEA.Population(), pps.MOEA.Archive())
}

func (pps PPS) plotMetric() {
	prob := pps.CMOP.Name() + "." + pps.MOEA.CHM().Name()
	path := "graphics/metric/" + prob
	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatal(err)
	}
	plotter := plotter.Plotter2D{
		Title:    prob + " Metric",
		LabelX:   "Generation",
		LabelY:   "IGD",
		Filename: path + "/" + prob + ".png",
	}
	plotter.PlotMetric(pps.MetricData, pps.SwitchPoint)
}

func (pps *PPS) printData(gen int) {
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fmt.Println(formatted, ",", gen, ",", pps.Stage(), ",", pps.MOEA.MaxViolation(), ",", pps.MOEA.FeasibleRatio(), ",", pps.MOEA.CHM().Threshold(gen), ",", pps.Performance())
}
