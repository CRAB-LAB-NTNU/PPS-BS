package pps

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"

	"github.com/CRAB-LAB-NTNU/PPS-BS/configs"
	"github.com/CRAB-LAB-NTNU/PPS-BS/plotter"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

// PPS is a struct describing the contents of the Push & Pull Framework
type PPS struct {
	cmop                                   types.CMOP
	moea                                   types.MOEA
	stages                                 []types.Stage
	stage                                  int
	idealPoints, nadirPoints, paretoPoints [][]float64
	SwitchPoint                            int
	MetricData                             []float64
	export                                 configs.Export
	Result                                 types.Results
}

func NewPPS(cmop types.CMOP, moea types.MOEA, stages []types.Stage, export configs.Export) PPS {
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

	pps.idealPoints = arrays.Zeros2DFloat64(pps.moea.MaxFuncEvals(), pps.cmop.NumberOfObjectives())
	pps.nadirPoints = arrays.Zeros2DFloat64(pps.moea.MaxFuncEvals(), pps.cmop.NumberOfObjectives())
	pps.stage = 0
	if points, err := plotter.ParseDatFile("arraydata/pf_data/" + pps.cmop.Name() + ".dat"); err == nil {
		pps.paretoPoints = points
	} else {
		fmt.Println("ERROR", err)
	}

}

// Run performs a run of the PPS framework
func (pps *PPS) Run() float64 {
	gen := 0
	for pps.moea.FunctionEvaluations() < pps.moea.MaxFuncEvals() {

		pps.setIdealAndNadir(gen)

		if pps.changeStage(gen) {
			pps.nextStage()
			pps.initStage()
		}
		pps.moea.Evolve(pps.currentStage())
		pps.printData(gen)

		if pps.export.ExportVideo {
			pps.plot(gen)
		}
		if pps.export.PlotEval {
			pps.MetricData = append(pps.MetricData, pps.Performance())
		}
		gen++
	}
	if pps.export.ExportVideo {
		pps.ExportVideo()
	}
	if pps.export.PlotEval {
		pps.plotMetric()
	}

	return pps.Performance()
}

func (pps *PPS) CMOP() types.CMOP {
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

func (pps PPS) ExportVideo() {
	prob := pps.cmop.Name() + "." + pps.moea.CHM().Name()
	path := "graphics/vids/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("path eksisterer ikke, produserer.")
		os.MkdirAll(path, 0755)
	}

	if err := os.Remove(path + prob + ".mp4"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Removed old file.")
	}

	cmd := exec.Command("ffmpeg", "-framerate", "20", "-i", "./graphics/gif/"+prob+"/%00d.png", "./graphics/vids/"+prob+".mp4")
	if err := cmd.Run(); err != nil {
		fmt.Println("Feil ved laging av video")
		log.Fatal(err)
	}
}

func (pps PPS) Performance() float64 {
	return pps.export.Metric(pps.moea.Archive(), pps.paretoPoints)
}
