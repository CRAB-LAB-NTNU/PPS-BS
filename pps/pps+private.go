package pps

import (
	"fmt"
	"time"

	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/plotter"
	"github.com/CRAB-LAB-NTNU/PPS-BS/stages"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func (pps *PPS) setIdealAndNadir() {
	ip, np := biooperators.CalculateNadirAndIdealPoints(pps.moea.Population())
	//Append in stead of initialising arrays at a fixed size to
	pps.idealPoints = append(pps.idealPoints, ip)
	pps.nadirPoints = append(pps.nadirPoints, np)
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
	push.Update(generation, pps.moea.Population())
	pps.stages[pps.stage] = push
	return pps.currentStage().IsOver()
}

func (pps *PPS) initStage() {
	if pps.currentStage().Type() == types.Pull {
		pps.initPull()
	}
	//pps.MOEA.CHM().Initialise(pps.MOEA.Generation(), pps.MOEA.MaxViolation())
}

func (pps *PPS) initPull() {
	pps.moea.CHM().Initialise(pps.generation, pps.moea.MaxViolation())
}

func (pps PPS) currentStage() types.Stage {
	return pps.stages[pps.stage]
}

func (pps *PPS) nextStage() {
	pps.stage++
	//If there are no feasible solutions we simply skip the boundary stage
	if pps.currentStage().Type() == types.BinarySearch && len(pps.moea.Archive()) == 0 {
		pps.stage++
	}
}

func (pps *PPS) printData(gen int) {
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fmt.Println(formatted, ",", gen, ",", pps.Stage(), ",", pps.moea.MaxViolation(), ",", pps.moea.FeasibleRatio(), ",", pps.moea.CHM().Threshold(gen))
}

func (pps *PPS) setupPlotter() {
	pop := pps.moea.Population()
	arc := pps.moea.Archive()
	pps.plotter = plotter.PpsPlotter{
		Population: &pop,
		Archive:    &arc,
		Generation: &pps.generation,
		Stage:      &pps.stage,
		Config:     &pps.export,
		Cmop:       &pps.cmop,
		Stages:     pps.stages,
	}
}
