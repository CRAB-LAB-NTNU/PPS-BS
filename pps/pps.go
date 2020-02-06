package pps

import (
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

// PPS is a struct describing the contents of the Push & Pull Framework
type PPS struct {
	Cmop                     types.CMOP
	Moea                     types.MOEA
	Stage                    types.Stage
	IdealPoints, NadirPoints [][]float64
	RK, Delta, Epsilon       float64
	TC, L                    int
}

// Initialise initialises the PPS framework with a given CMOP, MOEA and CHM
func (pps *PPS) Initialise(cmop types.CMOP, moea types.MOEA, TC int, delta, epsilon float64) {
	pps.Cmop = cmop
	pps.Moea = moea

	pps.TC = TC
	pps.Delta = delta
	pps.Epsilon = epsilon

	pps.RK = 1.0
	pps.Stage = types.Push

	pps.Moea.InitialisePopulation()

	pps.IdealPoints = generateEmpty2DSliceFloat64(pps.Moea.MaxGeneration(), pps.Cmop.NumberOfObjectives)
	pps.NadirPoints = generateEmpty2DSliceFloat64(pps.Moea.MaxGeneration(), pps.Cmop.NumberOfObjectives)
}

func generateEmpty2DSliceFloat64(outerLength, innerLength int) [][]float64 {
	slice := make([][]float64, outerLength)
	for i := range slice {
		slice[i] = make([]float64, innerLength)
	}
	return slice
}

func (pps *PPS) run() {
	for generation := 0; generation < pps.Moea.MaxGeneration(); generation++ {

		// First we set the ideal and nadir points for this generation based on the current population
		pps.SetIdealAndNadirPoints(pps.Moea.Population(), generation)
		if generation >= pps.L {
			pps.CalculateMaxChange(generation)
		}
		// If the change in ideal or nadir points is lower than a user defined value then we change phases
		if pps.RK <= pps.Epsilon {
			pps.Stage = types.Pull
		}

		// We evolve the population one generation
		// How this is done will depend on the underlying moea and constraint-handling method
		pps.Moea.Evolve(pps.Stage)

	}
}

//SetIdealAndNadirPoints sets the ideal and nadir points for the given generation
//This is done by first setting the ideal and nadir values to either max or min based on if the problem is a maximisation or minimisation problem.
//Then for each objective the value of the individual is checked to see if either the nadir or ideal point for the generation can be updated
func (pps *PPS) SetIdealAndNadirPoints(population []types.Individual, generation int) {
	// First set ideal and nadir points to best possible or worst possible values
	// based on if the objective is aoptimisation objective or minimisation objective

	/*
		for objectiveNr, objective := range pps.Cmop.Objectives() {

			switch objective.Type {
			case Minimisation:
				pps.IdealPoints[generation][objectiveNr] = math.MaxFloat64
				pps.NadirPoints[generation][objectiveNr] = math.SmallestNonzeroFloat64
			case Maximisation:
				pps.IdealPoints[generation][objectiveNr] = math.SmallestNonzeroFloat64
				pps.NadirPoints[generation][objectiveNr] = math.MaxFloat64
			}

			//TODO: Might have to redo the structure of how objective values are calculated
			for _, individual := range pps.Moea.Population() {
				for objectiveNr, objective := range pps.Cmop.Objectives() {
					objectiveValue := objective.Function(individual)
					idealPoint := pps.IdealPoints[generation][objectiveNr]
					nadirPoint := pps.NadirPoints[generation][objectiveNr]

					switch objective.Type {
					case Minimisation:
						if objectiveValue < idealPoint {
							idealPoint = objectiveValue
						} else if objectiveValue > nadirPoint {
							nadirPoint = objectiveValue
						}
					case Maximisation:
						if objectiveValue > idealPoint {
							idealPoint = objectiveValue
						} else if objectiveValue < nadirPoint {
							nadirPoint = objectiveValue
						}
					}
				}
			}
		}
	*/
}

// CalculateMaxChange Calculates the max change in ideal or nadir points
// Loops through each objective for the generation and finds the larges change from a previous generation.
func (pps *PPS) CalculateMaxChange(generation int) {
	/*
		maxChangeIdeal, maxChangeNadir := math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64
		for objectiveNr := range pps.Cmop.Objectives() {
			currentIdealPoint := pps.IdealPoints[generation][objectiveNr]
			currentNadirPoint := pps.NadirPoints[generation][objectiveNr]
			intervalIdealPoint := pps.IdealPoints[generation-pps.L][objectiveNr]
			intervalNadirPoint := pps.NadirPoints[generation-pps.L][objectiveNr]

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
	*/
}
