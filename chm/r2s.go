package chm

import (
	"fmt"
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

// R2S struct defines the parameters needed by r2s and the methods available
type R2S struct {
	Z, ZMin float64
	//InitialDeltaIn, InitialDeltaOut float64
	DeltaIn, DeltaOut           []float64
	hasCheckedActiveConstraints bool
	ActiveConstraints           []bool
	Val                         float64
	Cs, FESc, FESmax, FESacd    int
}

func NewR2S(FESc, FESacd, cs int, val, zMin float64, constraintsCount, generations int) *R2S {
	return &R2S{
		DeltaIn:           make([]float64, generations),
		DeltaOut:          make([]float64, generations),
		ActiveConstraints: make([]bool, constraintsCount),
		FESmax:            generations,
		FESc:              FESc,
		FESacd:            FESacd,
		Cs:                cs,
		Val:               val,
		ZMin:              zMin,
	}
}

func (r2s *R2S) Name() string {
	return "R2S"
}

func (r2s *R2S) Threshold(gen int) float64 {
	return r2s.DeltaOut[gen]
}

func (r2s *R2S) Initialise(t int, maxviolation float64) {

	r2s.DeltaIn[0] = maxviolation
	r2s.DeltaIn[t] = maxviolation
	r2s.DeltaOut[0] = maxviolation
	r2s.DeltaOut[t] = maxviolation

	r2s.initializeZ()

}

/*
func (r2s *R2S) Initialise(generations int, feasibleRatio float64, population []types.Individual) {
	fmt.Println("Initialising R2S")
	r2s.DeltaIn, r2s.DeltaOut = make([]float64, generations), make([]float64, generations)
	r2s.ActiveConstraints = make([]bool, population[0].Fitness().ConstraintCount)
	//TODO Set at parameter in a better way
	//TODO At the end of push phase set these using the population instead
	r2s.initializeDeltaIn(2)
	r2s.initializeDeltaOut(feasibleRatio, population)
	r2s.initializeZ()
	r2s.updateZ()

}
*/

// InitializeDeltaIn is used to initialize deltaIn to the input parameter passed to the method
func (r2s *R2S) initializeDeltaIn(initialDeltaIn float64) {
	//TODO: See how what effect changing to calculate the max constraint violation of all minimum constraint violations has
	fmt.Println("DeltaIn[0]: ", initialDeltaIn)
	r2s.DeltaIn[0] = initialDeltaIn
}

// InitializeDeltaOut is used to set an initial value for deltaOut
func (r2s *R2S) initializeDeltaOut(feasibleRatio float64, population []types.Individual) {

	//If number of feasible solutions is below 20% we calculate using maxviolation of the 20 "best" individuals
	//If not it is set to 1.
	//TODO: evalute if there are better approaches with better synergy with PPS
	if feasibleRatio < 0.2 {
		rankedPopulation := biooperators.FastNonDominatedSort(population)
		bestIndividuals := make([]types.Individual, len(population)*1/5)
		i := 0
		for _, nonDominatingSet := range rankedPopulation {
			for _, individual := range nonDominatingSet {
				bestIndividuals[i] = individual
				i++
				if i == len(bestIndividuals) {
					break
				}
			}
			if i == len(bestIndividuals) {
				break
			}
		}
		sumConstraintViolation := 0.0
		for _, individual := range bestIndividuals {
			//fmt.Println(i, "\tTotal Constraint Violation: ", individual.Fitness().TotalViolation())
			sumConstraintViolation += individual.Fitness().TotalViolation()
		}
		r2s.DeltaOut[0] = sumConstraintViolation / float64(len(bestIndividuals))
	} else {
		r2s.DeltaOut[0] = 1
	}
	fmt.Println("DeltaOut[0]: ", r2s.DeltaOut[0])

}

// Update updates the deltaIn and deltaOut of r2s.
// Require that cfe is a float to allow the use of interface for constraint handling.
func (r2s *R2S) Update(t int, cfe float64) {
	if r2s.HasActiveConstraints() {
		r2s.updateDeltaIn(t, int(cfe))
		r2s.updateDeltaOut(t, int(cfe))
	}
}

// UpdateDeltaIn is used to update deltaIn for each generation
func (r2s *R2S) updateDeltaIn(t int, cfe int) {

	minDeltaIn := 0.002 * r2s.DeltaIn[0]

	p1 := float64(cfe)
	numerator := r2s.DeltaIn[0] - minDeltaIn
	denominator := float64(r2s.FESmax)
	calcDeltaIn := r2s.DeltaIn[0] - p1*(numerator/denominator)

	r2s.DeltaIn[t] = math.Max(minDeltaIn, calcDeltaIn)

	//fmt.Println("DeltaIn[", t, "]=", r2s.DeltaIn[t])
}

// UpdateDeltaOut is used to calculate a new value for deltaOut each generation
func (r2s *R2S) updateDeltaOut(t int, cfe int) {
	if cfe <= r2s.FESc {
		p1 := r2s.DeltaOut[0]
		numerator := float64(cfe)
		denominator := float64(r2s.FESc)

		p2 := 1 - (numerator / denominator)

		r2s.DeltaOut[t] = p1 * math.Pow(p2, r2s.Z)

	} else {
		r2s.DeltaOut[t] = 0.0
	}
	//fmt.Println("DeltaOut[", t, "]=", r2s.DeltaOut[t])

}

// InitializeZ initialises Z based on the initial deltaout value
func (r2s *R2S) initializeZ() {

	numerator := -5 - math.Log(r2s.DeltaOut[0])
	denominator := math.Log(0.05)

	r2s.Z = numerator / denominator
}

// UpdateZ Updates Z using the Zmin value and the current Z value
func (r2s *R2S) updateZ() {
	//r2s.Z = math.Max(r2s.Z, r2s.ZMin)
	r2s.Z = 0.3*r2s.Z + 0.7*r2s.ZMin
}

//ACD check for active constraints near an assumed optimal individual
func (r2s *R2S) ACD(iter, cfe int, fitness types.Fitness) {
	if r2s.hasCheckedActiveConstraints {
		return
	}
	r2s.hasCheckedActiveConstraints = true
	activeConstraints := make([]bool, fitness.ConstraintCount)
	fmt.Println("Generation:", iter, "\tUpdating active constraints!")
	for constraint, constraintVal := range fitness.ConstraintValues {
		if r2s.constraintIsActive(constraintVal) {
			activeConstraints[constraint] = true
		}
	}
	r2s.ActiveConstraints = activeConstraints

	fmt.Println("Constraint Vals:", fitness.ConstraintValues)
	fmt.Println("Active Constraints: ", r2s.ActiveConstraints)

}

func (r2s *R2S) constraintIsActive(constraintVal float64) bool {
	return math.Abs(constraintVal) <= r2s.Val
}

//HasActiveConstraints check if any of the constraints of the problem are seen as active
func (r2s R2S) HasActiveConstraints() bool {
	for _, val := range r2s.ActiveConstraints {
		if val {
			return true
		}
	}
	return false
}

func (r2s R2S) Violation(t int, fitness types.Fitness) float64 {
	var total float64

	if !r2s.HasActiveConstraints() {
		return fitness.TotalViolation()
	}

	for c, isActiveConstraint := range r2s.ActiveConstraints {

		if !isActiveConstraint {
			continue
		}

		l := r2s.l(t, fitness.ConstraintTypes[c], fitness.ConstraintValues[c])
		r := r2s.r(t, fitness.ConstraintTypes[c], fitness.ConstraintValues[c])

		if l >= 0 && l <= r2s.DeltaIn[t] && r >= 0 && r <= r2s.DeltaOut[t] {
			continue
		} else {
			total += math.Min(math.Abs(l), math.Abs(r))
		}

	}
	return math.Abs(total)
}

func (r2s R2S) l(t int, constraintType types.ConstraintType, constraintViolation float64) float64 {
	if constraintType == types.EqualsOrGreaterThanZero {
		return r2s.DeltaIn[t] - math.Max(0.0, constraintViolation)
	} else {
		return r2s.DeltaIn[t] - math.Max(0.0, math.Abs(constraintViolation))
	}
}

func (r2s R2S) r(t int, constraintType types.ConstraintType, constraintViolation float64) float64 {
	if constraintType == types.EqualsOrGreaterThanZero {
		return r2s.DeltaOut[t] - math.Max(0.0, math.Abs(constraintViolation))
	} else {
		return r2s.DeltaOut[t] - math.Max(0.0, constraintViolation)
	}
}
