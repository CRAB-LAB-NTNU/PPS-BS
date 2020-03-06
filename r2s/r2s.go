package r2s

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
	DeltaIn, DeltaOut        []float64
	ActiveConstraints        []bool
	val                      float64
	cs, FESacd, FESmax, FESc int
}

func (r2s *R2S) Initialize(generations int, zMin, initialDeltaIn, feasibleRatio float64, population []types.Individual) {
	fmt.Println("Initialising R2S")
	r2s.DeltaIn, r2s.DeltaOut = make([]float64, generations), make([]float64, generations)
	r2s.InitializeDeltaIn(initialDeltaIn)
	r2s.InitializeDeltaOut(feasibleRatio, population)
	r2s.InitializeZMin(zMin)
	r2s.InitializeZ()
	r2s.UpdateZ()

	//TODO: allow user to set these parameters
	r2s.cs = 25
	r2s.FESacd = 200000
	r2s.val = 0.01
	r2s.FESmax = 300000
	r2s.FESc = 210000

}

// InitializeDeltaIn is used to initialize deltaIn to the input parameter passed to the method
func (r2s *R2S) InitializeDeltaIn(initialDeltaIn float64) {
	//TODO: See how what effect changing to calculate the max constraint violation of all minimum constraint violations has
	fmt.Println("Setting DeltaIn[0]: ", initialDeltaIn)
	r2s.DeltaIn[0] = initialDeltaIn
}

// InitializeDeltaOut is used to set an initial value for deltaOut
func (r2s *R2S) InitializeDeltaOut(feasibleRatio float64, population []types.Individual) {
	fmt.Println("Setting initial Delta out")

	//If number of feasible solutions is below 20% we calculate using maxviolation of the 20 "best" individuals
	//If not it is set to 1.
	//TODO: evalute if there are better approaches with better synergy with PPS
	if feasibleRatio < 0.2 {
		fmt.Println("Feasible ratio below 20%")
		rankedPopulation := biooperators.FastNonDominatedSort(population)
		bestIndividuals := make([]types.Individual, len(population))
		for i := 0; i < len(bestIndividuals); i++ {
			for _, nonDominatingSet := range rankedPopulation {
				for _, individual := range nonDominatingSet {
					bestIndividuals[i] = individual
				}
			}
		}
		sumConstraintViolation := 0.0
		for _, individual := range bestIndividuals {
			sumConstraintViolation += individual.Fitness().TotalViolationAbsolute()
		}
		r2s.DeltaOut[0] = sumConstraintViolation
	} else {
		r2s.DeltaOut[0] = 1
	}
	fmt.Println("DeltaOut[0]: ", r2s.DeltaOut[0])

}

// UpdateDeltaIn is used to update deltaIn for each generation
func (r2s *R2S) UpdateDeltaIn(t, cfe int) {
	p1 := float64(cfe)
	numerator := r2s.DeltaIn[0] - 0.002*r2s.DeltaIn[0]
	denominator := float64(r2s.FESmax)

	r2s.DeltaIn[t] = r2s.DeltaIn[0] - p1*(numerator/denominator)
	fmt.Println("DeltaIn[", t, "]=", r2s.DeltaIn[t])
}

// UpdateDeltaOut is used to calculate a new value for deltaOut each generation
func (r2s *R2S) UpdateDeltaOut(t, cfe int) {

	if cfe <= r2s.FESc {
		p1 := r2s.DeltaOut[0]

		numerator := float64(cfe)
		denominator := float64(r2s.FESc)

		p2 := 1 - (numerator / denominator)

		r2s.DeltaOut[t] = p1 * math.Pow(p2, r2s.Z)
		fmt.Println("DeltaOut[", t, "]=", r2s.DeltaOut[t])

	} else {
		r2s.DeltaOut[t] = 0.0
		fmt.Println("DeltaOut[", t, "]=", r2s.DeltaOut[t])
	}

}

// InitializeZMin initialises Zmin to the input parameter
func (r2s *R2S) InitializeZMin(zMin float64) {
	r2s.ZMin = zMin
}

// InitializeZ initialises Z based on the initial deltaout value
func (r2s *R2S) InitializeZ() {

	numerator := -5 - math.Log(r2s.DeltaOut[0])
	denominator := math.Log(0.05)

	r2s.Z = numerator / denominator
}

// UpdateZ Updates Z using the Zmin value and the current Z value
func (r2s *R2S) UpdateZ() {
	r2s.Z = 0.3*r2s.Z + 0.7*r2s.ZMin
}

func (r2s *R2S) ACD(iter, cfe int, fitness types.Fitness) {
	activeConstraints := make([]bool, fitness.ConstraintCount)
	if iter%r2s.cs == 0 && cfe <= r2s.FESacd {
		fmt.Println("Updating active constraints!")
		for constraint, constraintVal := range fitness.ConstraintValues {
			if r2s.ConstraintIsActive(constraintVal) {
				activeConstraints[constraint] = true
			}
		}
		r2s.ActiveConstraints = activeConstraints
		fmt.Println("Active Constraints: ", r2s.ActiveConstraints)
	}

}

func (r2s *R2S) ConstraintIsActive(constraintVal float64) bool {
	return constraintVal <= r2s.val
}

func (r2s R2S) HasActiveConstraints() bool {
	for _, val := range r2s.ActiveConstraints {
		if val {
			return true
		}
	}
	return false
}

func (r2s R2S) ConstraintViolation(t int, fitness types.Fitness) float64 {
	var total float64

	if !r2s.HasActiveConstraints() {
		return fitness.TotalViolationAbsolute()
	}

	for c, isActiveConstraint := range r2s.ActiveConstraints {
		if !isActiveConstraint {
			continue
		}
		l := r2s.l(t, fitness.ConstraintValues[c])
		r := r2s.r(t, fitness.ConstraintValues[c])

		if l <= 0 && r <= 0 {
			return 0
		}

		total += math.Max(0, math.Max(l, r))
	}
	return total
}

func (r2s R2S) l(t int, constraintViolation float64) float64 {
	return math.Abs(r2s.DeltaIn[t] - constraintViolation)
}

func (r2s R2S) r(t int, constraintViolation float64) float64 {
	return math.Abs(constraintViolation) - r2s.DeltaOut[t]
}
