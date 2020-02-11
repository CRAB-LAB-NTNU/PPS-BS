package pps

import (
	"fmt"
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

// PPS is a struct describing the contents of the Push & Pull Framework
type PPS struct {
	Cmop                                                           types.CMOP
	Moea                                                           types.MOEA
	Stage                                                          types.Stage
	idealPoints, nadirPoints                                       [][]float64
	RK, Delta, Epsilon                                             float64
	SearchingPreference, ConstraintRelaxation, RelaxationReduction float64
	TC, L                                                          int
	improvedEpsilon                                                []float64
	switchPopulation                                               []types.Individual
}

func (pps PPS) SwitchPopulation() []types.Individual {
	return pps.switchPopulation
}

// Initialise initialises the PPS framework with a given CMOP, MOEA and CHM
func (pps *PPS) Initialise() {
	pps.improvedEpsilon = make([]float64, pps.Moea.MaxGeneration())
	pps.Moea.Initialise()
	pps.idealPoints = generateEmpty2DSliceFloat64(pps.Moea.MaxGeneration(), pps.Cmop.NumberOfObjectives)
	pps.nadirPoints = generateEmpty2DSliceFloat64(pps.Moea.MaxGeneration(), pps.Cmop.NumberOfObjectives)
}

func generateEmpty2DSliceFloat64(outerLength, innerLength int) [][]float64 {
	slice := make([][]float64, outerLength)
	for i := range slice {
		slice[i] = make([]float64, innerLength)
	}
	return slice
}

func (pps *PPS) Run() {
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
			if pps.RK <= pps.Epsilon && pps.Stage == types.Push {
				fmt.Println("Switching stage")

				pps.Stage = types.Pull

				pps.switchPopulation = make([]types.Individual, len(pps.Moea.Population()))
				copy(pps.switchPopulation, pps.Moea.Population())

				pps.improvedEpsilon[generation], pps.improvedEpsilon[0] = pps.Moea.MaxViolation(), pps.Moea.MaxViolation()
			}

			if pps.Stage == types.Pull {
				pps.updateEpsilon(generation)
			}
		} else {
			pps.improvedEpsilon[generation] = 0
		}

		// We evolve the population one generation
		// How this is done will depend on the underlying moea and constraint-handling method
		pps.Moea.Evolve(pps.Stage, pps.improvedEpsilon)

	}
}

func (pps *PPS) updateEpsilon(k int) {

	if pps.RK < pps.SearchingPreference {
		pps.improvedEpsilon[k] = (1 - pps.ConstraintRelaxation) * pps.improvedEpsilon[k-1]
	} else {
		pps.improvedEpsilon[k] = pps.improvedEpsilon[0] * math.Pow((1-(float64(k)/float64(pps.Moea.MaxGeneration()))), pps.RelaxationReduction)
	}

}

// CalculateMaxChange Calculates the max change in ideal or nadir points
// Loops through each objective for the generation and finds the larges change from a previous generation.
func (pps *PPS) CalculateMaxChange(generation int) {
	maxChangeIdeal, maxChangeNadir := math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64
	for objectiveNr := 0; objectiveNr < pps.Cmop.NumberOfObjectives; objectiveNr++ {
		currentIdealPoint := pps.idealPoints[generation][objectiveNr]
		currentNadirPoint := pps.nadirPoints[generation][objectiveNr]
		intervalIdealPoint := pps.idealPoints[generation-pps.L][objectiveNr]
		intervalNadirPoint := pps.nadirPoints[generation-pps.L][objectiveNr]

		changeIdeal := math.Abs(currentIdealPoint-intervalIdealPoint) / math.Max(math.Abs(intervalIdealPoint), pps.Delta)
		changeNadir := math.Abs(currentNadirPoint-intervalNadirPoint) / math.Max(math.Abs(intervalNadirPoint), pps.Delta)
		if changeIdeal > maxChangeIdeal {
			maxChangeIdeal = changeIdeal
		}
		if changeNadir > maxChangeNadir {
			maxChangeNadir = changeNadir
		}
	}
	pps.RK = math.Max(maxChangeIdeal, maxChangeNadir)
}
