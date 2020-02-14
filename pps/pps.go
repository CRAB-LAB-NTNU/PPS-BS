package pps

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"

	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
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
	TC, L                                                          int
	improvedEpsilon                                                []float64
	paretoPoints                                                   [][]float64
}

// Initialise initialises the PPS framework with a given CMOP, MOEA and CHM
func (pps *PPS) Initialise() {
	pps.improvedEpsilon = make([]float64, pps.Moea.MaxGeneration())
	pps.Moea.Initialise()
	pps.idealPoints = generateEmpty2DSliceFloat64(pps.Moea.MaxGeneration(), pps.Cmop.NumberOfObjectives)
	pps.nadirPoints = generateEmpty2DSliceFloat64(pps.Moea.MaxGeneration(), pps.Cmop.NumberOfObjectives)
	pps.rk = 1.0
	pps.stage = types.Push
	pps.paretoPoints = plotter.ParseDatFile("arraydata/pf_data/" + pps.Cmop.Name() + ".dat")
}

func generateEmpty2DSliceFloat64(outerLength, innerLength int) [][]float64 {
	slice := make([][]float64, outerLength)
	for i := range slice {
		slice[i] = make([]float64, innerLength)
	}
	return slice
}

func (pps *PPS) Run() {
	fmt.Println("Solving", pps.Cmop.Name())
	for generation := 0; pps.Moea.FunctionEvaluations() < pps.Moea.MaxGeneration(); generation++ {

		// First we set the ideal and nadir points for this generation based on the current population
		ip, np := biooperators.CalculateNadirAndIdealPoints(pps.Moea.Population())
		pps.idealPoints[generation] = ip
		pps.nadirPoints[generation] = np

		if generation >= pps.L {
			pps.CalculateMaxChange(generation)
		}
		// If the change in ideal or nadir points is lower than a user defined value then we change phases
		if generation <= pps.TC {
			if pps.rk <= pps.Epsilon && pps.stage != types.Pull {
				fmt.Println("Switching stage")
				pps.stage = types.Pull
				pps.improvedEpsilon[generation], pps.improvedEpsilon[0] = pps.Moea.MaxViolation(), pps.Moea.MaxViolation()
			}

			if pps.stage == types.Pull {
				pps.updateEpsilon(generation)
			}
		} else {
			pps.improvedEpsilon[generation] = 0
		}

		// We evolve the population one generation
		// How this is done will depend on the underlying moea and constraint-handling method
		pps.Moea.Evolve(pps.stage, pps.improvedEpsilon)
		pps.plot(generation)
	}
}

func (pps *PPS) updateEpsilon(k int) {

	if pps.Moea.FeasibleRatio() < pps.SearchingPreference {
		pps.improvedEpsilon[k] = (1 - pps.ConstraintRelaxation) * pps.improvedEpsilon[k-1]
	} else {
		pps.improvedEpsilon[k] = pps.improvedEpsilon[0] * math.Pow((1-(float64(k)/float64(pps.Moea.MaxGeneration()))), pps.RelaxationReduction)
	}

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
	for i := 0; i < pps.Cmop.NumberOfObjectives; i++ {
		cur := points[k][i]
		offset := points[k-pps.L][i]
		dist := math.Abs(cur - offset)
		if calc := dist / math.Max(math.Abs(offset), pps.Delta); calc > m {
			m = calc
		}
	}
	return m
}

func (pps PPS) plot(generation int) {

	gen := strconv.Itoa(generation)
	eps := strconv.FormatFloat(pps.improvedEpsilon[generation], 'E', -1, 64)
	prob := pps.Cmop.Name()
	var stage string
	if pps.stage == types.Push {
		stage = "PUSH"
	} else {
		stage = "PULL"
	}

	path := "graphics/gif/" + prob

	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatal(err)
	}

	plotter := plotter.Plotter2D{
		Title:    prob + " Stage: " + stage + " gen: " + gen + " eps: " + eps,
		LabelX:   "f1",
		LabelY:   "f2",
		Min:      0,
		Max:      3,
		Filename: path + "/" + gen + ".png",
		Solution: pps.paretoPoints,
		Extremes: [][]float64{pps.idealPoints[generation], pps.nadirPoints[generation], pps.Moea.Ideal()},
	}
	plotter.Plot(pps.Moea.Population(), pps.Moea.Archive())
}

func (pps PPS) ExportVideo() {
	prob := pps.Cmop.Name()
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
