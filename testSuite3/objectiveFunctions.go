package testSuite3

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
	return (1 + inner2(x)) * x[0]
}

func objective4(x types.Genotype) float64 {
	return (1 + inner2(x)) * (1 - math.Pow(x[0], 2))
}

func objective5(x types.Genotype) float64 {
	return (1 + inner3(x)) * math.Cos((math.Pi*x[0])/2)
}

func objective6(x types.Genotype) float64 {
	return (1 + inner3(x)) * math.Sin((math.Pi*x[0])/2)
}

func objective7(x types.Genotype) float64 {
	return (1 + inner4(x)) * x[0]
}

func objective8(x types.Genotype) float64 {
	a := (1 + inner4(x))
	b := 1 - math.Pow(x[0], 0.5)*math.Pow(math.Cos(2*math.Pi*x[0]), 2)
	return a * b
}
