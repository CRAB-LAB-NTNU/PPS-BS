package pps

import (
	"github.com/CRAB-LAB-NTNU/PPS-BS/configs"
	"github.com/CRAB-LAB-NTNU/PPS-BS/plotter"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

// PPS is a struct describing the contents of the Push & Pull Framework
type PPS struct {
	cmop                                   types.Cmop
	moea                                   types.MOEA
	stages                                 []types.Stage
	stage, generation                      int
	idealPoints, nadirPoints, paretoPoints [][]float64
	export                                 configs.Export
	plotter                                plotter.PpsPlotter
}

func NewPPS(cmop types.Cmop, moea types.MOEA, stages []types.Stage, export configs.Export) PPS {
	pps := PPS{
		cmop:   cmop,
		moea:   moea,
		stages: stages,
		export: export,
	}

	pps.Initialise()
	return pps

}

// Initialise initialises the PPS framework with a given CMOP, MOEA and CHM
func (pps *PPS) Initialise() {

	pps.stage = 0

}

// Run performs a run of the PPS framework
func (pps *PPS) Run() []types.Individual {
	if pps.export.ExportVideo {
		pps.setupPlotter()
	}
	for pps.moea.FunctionEvaluations() < pps.moea.MaxFuncEvals() {

		pps.setIdealAndNadir()

		if pps.changeStage(pps.generation) {
			pps.nextStage()
			pps.initStage()
			pop := pps.moea.Population()
			pps.plotter.Population = &pop
		}
		pps.moea.Evolve(pps.currentStage())

		if pps.export.PrintGeneration {
			pps.printData(pps.generation)
		}

		if pps.export.ExportVideo {
			arc := pps.moea.Archive()
			pps.plotter.Archive = &arc
			pps.plotter.PlotFrame()
		}

		pps.generation++
	}
	if pps.export.ExportVideo {
		pps.plotter.ExportVideo()
	}
	/* TODO needs to be implemented.
	if pps.export.PlotEval {
		pps.plotter.PlotMetric()
	}
	*/
	return pps.moea.Archive()
}

func (pps *PPS) CMOP() types.Cmop {
	return pps.cmop
}
func (pps *PPS) Stages() []string {
	var stages []string
	for _, stage := range pps.stages {
		stages = append(stages, stage.Name())
	}
	return stages
}

func (pps *PPS) MOEA() types.MOEA {
	return pps.moea
}

func (pps PPS) Stage() string {
	return pps.currentStage().Name()
}
