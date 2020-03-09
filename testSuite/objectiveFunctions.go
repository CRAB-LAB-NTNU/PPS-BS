package testSuite

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func Objective1(x types.Genotype) float64 {
	return x[0] + Inner1(x)
}

func Objective2(x types.Genotype) float64 {
	return 1 - math.Pow(x[0], 2) + Inner2(x)
}

func Objective3(x types.Genotype) float64 {
	return 1 - math.Sqrt(x[0]) + Inner2(x)
}

func Objective4(x types.Genotype) float64 {
	return x[0] + 10*Inner3(x) + 0.7057
}

func Objective5(x types.Genotype) float64 {
	return 1 - math.Sqrt(x[0]) + 10*Inner4(x) + 0.7057
}

func Objective6(x types.Genotype) float64 {
	return 1 - math.Pow(x[0], 2) + 10*Inner4(x) + 0.7057
}

func Objective7(x types.Genotype) float64 {
	return 1.7057 * x[0] * (10*Inner3(x) + 1)
}

func Objective8(x types.Genotype) float64 {
	return 1.7057 * (1 - math.Pow(x[0], 2)) * (10*Inner4(x) + 1)
}

func Objective9(x types.Genotype) float64 {
	return 1.7057 * (1 - math.Sqrt(x[0])) * (10*Inner4(x) + 1)
}

func Objective10(x types.Genotype) float64 {
	return (1.7057 + Inner5(x)) * math.Cos(0.5*math.Pi*x[0]) * math.Cos(0.5*math.Pi*x[1])
}

func Objective11(x types.Genotype) float64 {
	return (1.7057 + Inner5(x)) * math.Cos(0.5*math.Pi*x[0]) * math.Sin(0.5*math.Pi*x[1])
}

func Objective12(x types.Genotype) float64 {
	return (1.7057 + Inner5(x)) * math.Sin(0.5*math.Pi*x[0])
}
