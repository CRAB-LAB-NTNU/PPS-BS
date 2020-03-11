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
	CMOP                                   types.CMOP
	MOEA                                   types.MOEA
	Stages                                 []types.Stage
	stage                                  int
	idealPoints, nadirPoints, paretoPoints [][]float64
	SwitchPoint                            int
	MetricData                             []float64
	Export                                 configs.Export
	Result                                 types.Results
}

func (pps *PPS) Reset() {
	pps.MOEA.Reset()
	pps.Initialise()
}

// Initialise initialises the PPS framework with a given CMOP, MOEA and CHM
func (pps *PPS) Initialise() {

	pps.idealPoints = arrays.Zeros2DFloat64(pps.MOEA.MaxFuncEvals(), pps.CMOP.NumberOfObjectives())
	pps.nadirPoints = arrays.Zeros2DFloat64(pps.MOEA.MaxFuncEvals(), pps.CMOP.NumberOfObjectives())
	pps.stage = 0
	if points, err := plotter.ParseDatFile("arraydata/pf_data/" + pps.CMOP.Name() + ".dat"); err == nil {
		pps.paretoPoints = points
	} else {
		fmt.Println("ERROR", err)
	}

}

func (pps *PPS) Run() float64 {
	for generation := 0; pps.MOEA.FunctionEvaluations() < pps.MOEA.MaxFuncEvals(); generation++ {

		pps.setIdealAndNadir(generation)

		if pps.changeStage(generation) {
			pps.nextStage()
			pps.initStage()
		}

		pps.printData(generation)

		pps.MOEA.Evolve(pps.currentStage().Stage())

		if pps.Export.ExportVideo {
			pps.plot(generation)
		}
		if pps.Export.PlotEval {
			pps.MetricData = append(pps.MetricData, pps.Performance())
		}
	}
	if pps.Export.ExportVideo {
		pps.ExportVideo()
	}
	if pps.Export.PlotEval {
		pps.plotMetric()
	}

	return pps.Performance()
}

func (pps PPS) RunTest() {
	for i := 0; i < pps.Export.Runs; i++ {
		fmt.Println(pps.CMOP.Name(), "RUN:", i+1)
		pps.Result.Add(pps.Run())
		pps.Reset()
	}
	fmt.Println("PROBLEM:", pps.CMOP.Name())
	fmt.Println("Stages:", pps.Stages)
	fmt.Println("Constraint method:", pps.MOEA.GetCHM().Name())
	fmt.Println("MEAN:", pps.Result.Mean())
	fmt.Println("VAR:", pps.Result.Variance())
	fmt.Println("STD:", pps.Result.StandardDeviation())
	fmt.Println()
}

func (pps PPS) Stage() string {
	return pps.currentStage().Name()
}

func (pps PPS) ExportVideo() {
	prob := pps.CMOP.Name() + "." + pps.MOEA.GetCHM().Name()
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
	return pps.Export.Metric(pps.MOEA.Archive(), pps.paretoPoints)
}
