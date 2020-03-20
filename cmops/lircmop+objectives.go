package cmops

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func lirobjective1(x types.Genotype) float64 {
	return x[0] + lirinner1(x)
}

func lirobjective2(x types.Genotype) float64 {
	return 1 - math.Pow(x[0], 2) + lirinner2(x)
}

func lirobjective3(x types.Genotype) float64 {
	return 1 - math.Sqrt(x[0]) + lirinner2(x)
}

func lirobjective4(x types.Genotype) float64 {
	return x[0] + 10*lirinner3(x) + 0.7057
}

func lirobjective5(x types.Genotype) float64 {
	return 1 - math.Sqrt(x[0]) + 10*lirinner4(x) + 0.7057
}

func lirobjective6(x types.Genotype) float64 {
	return 1 - math.Pow(x[0], 2) + 10*lirinner4(x) + 0.7057
}

func lirobjective7(x types.Genotype) float64 {
	return 1.7057 * x[0] * (10*lirinner3(x) + 1)
}

func lirobjective8(x types.Genotype) float64 {
	return 1.7057 * (1 - math.Pow(x[0], 2)) * (10*lirinner4(x) + 1)
}

func lirobjective9(x types.Genotype) float64 {
	return 1.7057 * (1 - math.Sqrt(x[0])) * (10*lirinner4(x) + 1)
}

func lirobjective10(x types.Genotype) float64 {
	return (1.7057 + lirinner5(x)) * math.Cos(0.5*math.Pi*x[0]) * math.Cos(0.5*math.Pi*x[1])
}

func lirobjective11(x types.Genotype) float64 {
	return (1.7057 + lirinner5(x)) * math.Cos(0.5*math.Pi*x[0]) * math.Sin(0.5*math.Pi*x[1])
}

func lirobjective12(x types.Genotype) float64 {
	return (1.7057 + lirinner5(x)) * math.Sin(0.5*math.Pi*x[0])
}
