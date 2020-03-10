package chm

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

//Epsilon is used for Epsilon constraint handling
type Epsilon struct {
	Cp        float64
	Tc        int
	threshold []float64
}

func NewE(cp float64, tc, maxGenerations int) *Epsilon {
	return &Epsilon{
		Cp:        cp,
		Tc:        tc,
		threshold: make([]float64, maxGenerations),
	}
}

func (e Epsilon) Name() string {
	return "Epsilon"
}

func (e Epsilon) Threshold(gen int) float64 {
	return e.threshold[gen]
}

//Update updates the epsilon value used for the generation (k)
func (e *Epsilon) Update(t int) {
	if e.Tc < t {
		e.threshold[t] = 0
	} else {
		p1 := e.threshold[0]

		numerator := float64(t)
		denominator := float64(e.Tc)
		p2 := math.Pow((1 - numerator/denominator), e.Cp)

		e.threshold[t] = p1 * p2
	}
}

// Violation calculates the constraint violation in regards to the threshold set by the CHM.
// If the violation is below the threshold, 0 is returned.
func (e Epsilon) Violation(t int, f types.Fitness) float64 {
	constraintViolation := f.TotalViolation()

	if constraintViolation <= e.threshold[t] {
		return 0
	}
	return constraintViolation
}
