package pps

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/configs"
	"github.com/CRAB-LAB-NTNU/PPS-BS/metrics"
	"github.com/CRAB-LAB-NTNU/PPS-BS/plotter"
	"github.com/CRAB-LAB-NTNU/PPS-BS/sweeper"
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
	sweeper                                sweeper.Sweeper
	popResults, arcResults                 metrics.Results
	HVCoefficient                          float64
}

func NewPPS(cmop types.Cmop, moea types.MOEA, stages []types.Stage, sweeper sweeper.Sweeper, hvCoefficient float64, export configs.Export) PPS {
	pps := PPS{
		cmop:          cmop,
		moea:          moea,
		stages:        stages,
		sweeper:       sweeper,
		HVCoefficient: hvCoefficient,
		export:        export,
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
	if pps.sweeper.Sweep() {
		pps.popResults = metrics.Results{
			ParetoFront:          pps.cmop.TrueParetoFront(),
			HyperVolumeReference: metrics.HVReferenceNadirTimes(pps.HVCoefficient, pps.cmop),
		}
		pps.arcResults = metrics.Results{
			ParetoFront:          pps.cmop.TrueParetoFront(),
			HyperVolumeReference: metrics.HVReferenceNadirTimes(pps.HVCoefficient, pps.cmop),
		}
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

		if pps.sweeper.Sweep() {
			pps.popResults.Add(pps.moea.Population())
			pps.arcResults.Add(pps.moea.Archive())
			var values []interface{}
			values = append(values, pps.generation)
			if pps.sweeper.Phase() {
				values = append(values, pps.stages[pps.stage].Name())
			}
			if pps.sweeper.FR() {
				values = append(values, pps.moea.FeasibleRatio())
			}
			if pps.sweeper.CD() {
				popdist := biooperators.CrowdingDistance(pps.moea.Population())
				var s float64
				for _, val := range popdist {
					if val == math.MaxFloat64 {
						continue
					}
					s += val
				}
				values = append(values, s)
			}
			if pps.sweeper.IGD() {
				values = append(values, pps.popResults.IGD.Last())
			}
			if pps.sweeper.HV() {
				values = append(values, pps.popResults.HV.Last())
			}
			if pps.sweeper.ArchiveIGD() {
				values = append(values, pps.arcResults.IGD.Last())
			}
			if pps.sweeper.ArchiveHV() {
				values = append(values, pps.arcResults.HV.Last())
			}

			pps.sweeper.WriteLine(values)
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

func (pps PPS) Sweeper() sweeper.Sweeper {
	return pps.sweeper
}
