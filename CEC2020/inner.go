package CEC2020

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func g1(x types.Genotype) float64 {
	var s float64
	n := len(x)
	for i := 1; i < n; i++ {
		floatN, floatI := float64(n), float64(i)
		z := math.Pow(x[i], floatN-2)
		fraction := floatI / (2 * floatN)
		inner := math.Pow(z-0.5-fraction, 2)
		s += 1 - math.Exp(-10*inner)
	}
	return 1 + s
}
