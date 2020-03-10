package pps

import (
	"fmt"
	"math"
	"time"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/configs"
	"github.com/CRAB-LAB-NTNU/PPS-BS/plotter"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

// PPS is a struct describing the contents of the Push & Pull Framework
type PPS struct {
	Cmop                                                           types.CMOP
	Moea                                                           types.MOEA
	stage                                                          types.Stage
	idealPoints, nadirPoints                                       [][]float64
	rk, Delta, Epsilon                                             float64
	SearchingPreference, ConstraintRelaxation, RelaxationReduction float64
	TC, L, SwitchPoint, generation                                 int
	improvedEpsilon, MetricData                                    []float64
	Config                                                         configs.PPS
	Result                                                         types.Results
	DoBoundary                                                     bool
	ConstraintUpdateMethod                                         types.ConstraintMethod
	plotter                                                        plotter.PpsPlotter
}

// Initialise initialises the PPS framework with a given CMOP, MOEA and CHM
func (pps *PPS) Initialise() {
	pps.Moea.Initialise(pps.Cmop)

	genMax := pps.Moea.MaxEvaluations() / len(pps.Moea.Population())

	pps.improvedEpsilon = make([]float64, genMax)
	pps.idealPoints = arrays.GenerateEmpty2DSliceFloat64(genMax, pps.Cmop.ObjectiveCount)
	pps.nadirPoints = arrays.GenerateEmpty2DSliceFloat64(genMax, pps.Cmop.ObjectiveCount)
	pps.rk = 1.0
	pps.stage = types.Push
	if pps.Config.ExportVideo {
		pop := pps.Moea.Population()
		arc := pps.Moea.Archive()
		pps.plotter = plotter.PpsPlotter{
			Population: &pop,
			Archive:    &arc,
			Ideal:      &pps.idealPoints,
			Nadir:      &pps.nadirPoints,
			Epsilon:    &pps.improvedEpsilon,
			Generation: &pps.generation,
			Stage:      &pps.stage,
			Config:     &pps.Config,
			Cmop:       &pps.Cmop,
		}
	}

}

func (pps *PPS) Run() float64 {
	for pps.Moea.FunctionEvaluations() < pps.Moea.MaxEvaluations() {

		// First we set the ideal and nadir points for this generation based on the current population
		ip, np := biooperators.CalculateNadirAndIdealPoints(pps.Moea.Population())
		pps.idealPoints[pps.generation] = ip
		pps.nadirPoints[pps.generation] = np

		if pps.generation >= pps.L {
			pps.CalculateMaxChange(pps.generation)
		}
		// If the change in ideal or nadir points is lower than a user defined value then we change phases
		if pps.generation <= pps.TC {
			if pps.rk <= pps.Epsilon && pps.stage < types.BorderSearch {
				if pps.DoBoundary {
					pps.stage = types.BorderSearch
				} else {
					pps.stage = types.Pull
					pps.improvedEpsilon[pps.generation], pps.improvedEpsilon[0] = pps.Moea.MaxViolation(), pps.Moea.MaxViolation()
				}
			} else if pps.stage < types.Pull && pps.Moea.BinaryDone() {
				pop := pps.Moea.Population()
				pps.plotter.Population = &pop
				pps.stage = types.Pull
				pps.improvedEpsilon[pps.generation], pps.improvedEpsilon[0] = pps.Moea.MaxViolation(), pps.Moea.MaxViolation()
			} else if pps.stage == types.Pull {

				switch pps.ConstraintUpdateMethod {
				case types.ImprovedEpsilon:
					pps.UpdateImprovedEpsilon()
				case types.Epsilon:
					pps.UpdateEpsilon()
				}

			}
		} else {
			pps.improvedEpsilon[pps.generation] = 0
		}
		pps.printData()
		// We evolve the population one generation
		// How this is done will depend on the underlying moea and constraint-handling method

		pps.Moea.Evolve(pps.stage, pps.improvedEpsilon)

		if pps.Config.ExportVideo {
			// Hacky way of fixing archive reference
			arc := pps.Moea.Archive()
			pps.plotter.Archive = &arc
			pps.plotter.PlotFrame()
		}
		if pps.Config.PlotEval {
			pps.MetricData = append(pps.MetricData, pps.Performance())
		}
		pps.generation++
	}
	if pps.Config.ExportVideo {
		pps.plotter.ExportVideo()

	}
	if pps.Config.PlotEval {
		//plotter.PlotMetric()
	}

	return pps.Performance()
}

func (pps PPS) RunTest() {
	for i := 0; i < pps.Config.Runs; i++ {
		pps.Initialise()
		fmt.Println(pps.Cmop.Name, "RUN:", i+1)
		pps.Result.Add(pps.Run())
	}
	fmt.Println("PROBLEM:", pps.Cmop.Name)
	fmt.Println("Boundary Search:", pps.DoBoundary)
	fmt.Println("Constraint method:", pps.ConstraintUpdateMethod)
	fmt.Println("MEAN:", pps.Result.Mean())
	fmt.Println("VAR:", pps.Result.Variance())
	fmt.Println("STD:", pps.Result.StandardDeviation())
	fmt.Println()
}

func (pps *PPS) UpdateImprovedEpsilon() {
	k := pps.generation
	if pps.Moea.FeasibleRatio() < pps.SearchingPreference {
		pps.improvedEpsilon[k] = (1 - pps.ConstraintRelaxation) * pps.improvedEpsilon[k-1]
	} else {
		pps.improvedEpsilon[k] = pps.improvedEpsilon[0] * math.Pow((1-(float64(k)/float64(pps.TC))), pps.RelaxationReduction)
	}

}

func (pps *PPS) UpdateEpsilon() {
	k := pps.generation
	pps.improvedEpsilon[k] = (1 - pps.ConstraintRelaxation) * pps.improvedEpsilon[k-1]
}

// CalculateMaxChange Calculates the max change in ideal or nadir points
// Loops through each objective for the generation and finds the larges change from a previous generation.
func (pps *PPS) CalculateMaxChange(generation int) {
	rz := pps.rx(generation, pps.idealPoints)
	rn := pps.rx(generation, pps.nadirPoints)
	pps.rk = math.Max(rz, rn)
}

func (pps PPS) rx(k int, points [][]float64) float64 {
	m := math.SmallestNonzeroFloat64
	for i := 0; i < pps.Cmop.ObjectiveCount; i++ {
		cur := points[k][i]
		offset := points[k-pps.L][i]
		dist := math.Abs(cur - offset)
		if calc := dist / math.Max(math.Abs(offset), pps.Delta); calc > m {
			m = calc
		}
	}
	return m
}

func (pps PPS) Stage() string {
	return []string{"Push", "Border Search", "Pull"}[pps.stage]
}

func (pps PPS) Performance() float64 {
	return pps.Config.Metric(pps.Moea.Archive(), pps.Cmop.TrueParetoFront())
}

func (pps *PPS) printData() {
	gen := pps.generation
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fmt.Println(formatted, ",", gen, ",", pps.Stage(), ",", pps.Moea.MaxViolation(), ",", pps.Moea.FeasibleRatio(), ",", pps.improvedEpsilon[gen] /*, ",", pps.Performance()*/)
}

/*
func (pps PPS) plot(generation int) {

	gen := strconv.Itoa(generation)
	eps := strconv.FormatFloat(pps.improvedEpsilon[generation], 'E', -1, 64)
	prob := pps.Cmop.Name
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
		Solution: pps.Cmop.TrueParetoFront(),
		Extremes: [][]float64{pps.idealPoints[generation], pps.nadirPoints[generation], pps.Moea.Ideal()},
	}
	plotter.Plot(pps.Moea.Population(), pps.Moea.Archive())
}

func (pps PPS) plotMetric() {
	prob := pps.Cmop.Name
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
	prob := pps.Cmop.Name
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
*/
