package testsuite2

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func objective1(x types.Genotype) float64 {
	return (1 + inner1(x)) * x[0]
}

func objective2(x types.Genotype) float64 {
	return (1 + inner1(x)) * (1 - math.Sqrt(x[0]))
}

func objective3(x types.Genotype) float64 {
	return (1 + inner1(x)) * (1 - x[0])
}

func objective4(x types.Genotype) float64 {
	return (1 + inner1(x)) * (1 - math.Pow(x[0], 2))
}
