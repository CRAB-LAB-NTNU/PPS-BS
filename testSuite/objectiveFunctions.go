package testSuite

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func objective1(x types.Genotype) float64 {
	return x[0] + inner1(x)
}

func objective2(x types.Genotype) float64 {
	return 1 - math.Pow(x[0], 2) + inner2(x)
}

func objective3(x types.Genotype) float64 {
	return 1 - math.Sqrt(x[0]) + inner2(x)
}

func objective4(x types.Genotype) float64 {
	return x[0] + 10*inner3(x) + 0.7057
}

func objective5(x types.Genotype) float64 {
	return 1 - math.Sqrt(x[0]) + 10*inner4(x) + 0.7057
}

func objective6(x types.Genotype) float64 {
	return 1 - math.Pow(x[0], 2) + 10*inner4(x) + 0.7057
}

func objective7(x types.Genotype) float64 {
	return 1.7057 * x[0] * (10*inner3(x) + 1)
}

func objective8(x types.Genotype) float64 {
	return 1.7057 * (1 - math.Pow(x[0], 2)) * (10*inner4(x) + 1)
}

func objective9(x types.Genotype) float64 {
	return 1.7057 * (1 - math.Sqrt(x[0])) * (10*inner4(x) + 1)
}

func objective10(x types.Genotype) float64 {
	return (1.7057 + inner5(x)) * math.Cos(0.5*math.Pi*x[0]) * math.Cos(0.5*math.Pi*x[1])
}

func objective11(x types.Genotype) float64 {
	return (1.7057 + inner5(x)) * math.Cos(0.5*math.Pi*x[0]) * math.Sin(0.5*math.Pi*x[1])
}

func objective12(x types.Genotype) float64 {
	return (1.7057 + inner5(x)) * math.Sin(0.5*math.Pi*x[0])
}
