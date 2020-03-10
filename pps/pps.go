package pps

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"

	"github.com/CRAB-LAB-NTNU/PPS-BS/configs"
	"github.com/CRAB-LAB-NTNU/PPS-BS/plotter"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

// PPS is a struct describing the contents of the Push & Pull Framework
type PPS struct {
	CMOP                                                           types.CMOP
	MOEA                                                           types.MOEA
	Stages                                                         []types.Stage
	stage                                                          int
	idealPoints, nadirPoints, paretoPoints                         [][]float64
	rk, Delta, Epsilon                                             float64
	SearchingPreference, ConstraintRelaxation, RelaxationReduction float64
	L, SwitchPoint                                                 int
	MetricData                                                     []float64
	Config                                                         configs.PPS
	Result                                                         types.Results
	DoBoundary                                                     bool
}

//TODO: this struct is becoming very large and not really that modular. should maybe look at ways to split things up more and make them general

func (pps *PPS) Reset() {
	pps.MOEA.Reset()
	pps.Initialise()
}

// Initialise initialises the PPS framework with a given CMOP, MOEA and CHM
func (pps *PPS) Initialise() {
	pps.MOEA.Initialise()

	pps.idealPoints = arrays.Zeros2DFloat64(pps.MOEA.MaxFuncEvals(), pps.CMOP.NumberOfObjectives())
	pps.nadirPoints = arrays.Zeros2DFloat64(pps.MOEA.MaxFuncEvals(), pps.CMOP.NumberOfObjectives())
	pps.rk = 1.0
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

		pps.calculateMaxChange(generation)

		//TODO check if change phase
		// need to handle first time in phase where constraints are considered. either here or in moead. maybe better in moead

		if pps.changeStage(generation) {
			pps.nextStage()
			//TODO Initialise boundary values
		}

		pps.printData(generation)

		// We evolve the population one generation
		// How this is done will depend on the underlying moea and constraint-handling method
		pps.MOEA.Evolve(pps.currentStage().Stage())

		if pps.Config.ExportVideo {
			pps.plot(generation)
		}
		if pps.Config.PlotEval {
			pps.MetricData = append(pps.MetricData, pps.Performance())
		}
	}
	if pps.Config.ExportVideo {
		pps.ExportVideo()
	}
	if pps.Config.PlotEval {
		pps.plotMetric()
	}

	return pps.Performance()
}

func (pps PPS) RunTest() {
	for i := 0; i < pps.Config.Runs; i++ {
		fmt.Println(pps.CMOP.Name(), "RUN:", i+1)
		pps.Result.Add(pps.Run())
		pps.Reset()
	}
	fmt.Println("PROBLEM:", pps.CMOP.Name())
	fmt.Println("Boundary Search:", pps.DoBoundary)
	fmt.Println("Constraint method:", pps.MOEA.GetCHM().Name())
	fmt.Println("MEAN:", pps.Result.Mean())
	fmt.Println("VAR:", pps.Result.Variance())
	fmt.Println("STD:", pps.Result.StandardDeviation())
	fmt.Println()
}

func (pps PPS) Stage() string {
	return []string{"Push", "Border Search", "Pull"}[pps.stage]
}

func (pps PPS) plot(generation int) {

	gen := strconv.Itoa(generation)
	eps := strconv.FormatFloat(pps.MOEA.GetCHM().Threshold(generation), 'E', -1, 64)
	prob := pps.CMOP.Name()
	stage := pps.Stage()

	path := "graphics/gif/" + prob

	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatal(err)
	}

	plotter := plotter.Plotter2D{
		Title:    prob + " Stage: " + stage + " gen: " + gen + " eps: " + eps,
		LabelX:   "f1",
		LabelY:   "f2",
		Min:      pps.Config.VideoMin,
		Max:      pps.Config.VideoMax,
		Filename: path + "/" + gen + ".png",
		Solution: pps.paretoPoints,
		Extremes: [][]float64{pps.idealPoints[generation], pps.nadirPoints[generation], pps.MOEA.Ideal()},
	}
	plotter.Plot(pps.MOEA.Population(), pps.MOEA.Archive())
}

func (pps PPS) plotMetric() {
	prob := pps.CMOP.Name()
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

func (pps PPS) ExportVideo() {
	prob := pps.CMOP.Name()
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
	return pps.Config.Metric(pps.MOEA.Archive(), pps.paretoPoints)
}

func (pps *PPS) printData(gen int) {
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fmt.Println(formatted, ",", gen, ",", pps.Stage(), ",", pps.MOEA.MaxViolation(), ",", pps.MOEA.FeasibleRatio(), ",", pps.MOEA.GetCHM().Threshold(gen), ",", pps.Performance())
}
