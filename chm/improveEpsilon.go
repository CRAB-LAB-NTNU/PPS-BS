package chm

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

//ImprovedEpsilon is used for Improved Epsilon Constraint Handling
type ImprovedEpsilon struct {
	tau, alpha, cp float64
	tc             int
	threshold      []float64
}

func NewIE(tau, alpha, cp float64, tc, maxGenerations int) *ImprovedEpsilon {

	return &ImprovedEpsilon{
		tau:       tau,
		alpha:     alpha,
		cp:        cp,
		tc:        tc,
		threshold: make([]float64, maxGenerations),
	}
}
func (ie *ImprovedEpsilon) Name() string {
	return "ImprovedEpsilon"
}

func (ie *ImprovedEpsilon) Threshold(gen int) float64 {
	return ie.threshold[gen]
}
func (ie *ImprovedEpsilon) Initialise(generation int, maxViolation float64) {
	ie.threshold[generation], ie.threshold[generation-1], ie.threshold[0] = maxViolation, maxViolation, maxViolation
}

//Update updates the epsilon value used for the generation (t) based on the feasibility ratio (rfk) of the current generation
func (ie *ImprovedEpsilon) Update(t int, rfk float64) {
	if ie.tc < t {

		ie.threshold[t] = 0
	} else if rfk < ie.alpha {

		p1 := (1 - ie.tau)
		p2 := ie.threshold[t-1]

		ie.threshold[t] = p1 * p2

	} else {
		p1 := ie.threshold[0]

		numerator := float64(t)
		denominator := float64(ie.tc)
		p2 := math.Pow((1 - numerator/denominator), ie.cp)

		ie.threshold[t] = p1 * p2
	}
}

// Violation calculates the constraint violation in regards to the threshold set by the CHM
// If the violation is below the threshold, 0 is returned
func (ie ImprovedEpsilon) Violation(t int, f types.Fitness) float64 {
	constraintViolation := f.TotalViolation()

	if constraintViolation <= ie.threshold[t] {
		return 0
	}

	return constraintViolation

}
