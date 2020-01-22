package testSuite

import (
	"math"
)

func cmop1F1(x []float64) float64 {
	return x[0] + cmop1G1(x)
}

func cmop1F2(x []float64) float64 {
	return 1 - math.Pow(x[0], 2) + cmop1G2(x)
}

func cmop1G1(x []float64) float64 {
	var s float64
	for i := 2; i < len(x); i += 2 {
		s += math.Pow(x[i]-math.Sin(0.5*math.Pi*x[0]), 2)
	}
	return s
}

func cmop1G2(x []float64) float64 {
	var s float64
	for i := 1; i < len(x); i += 2 {
		s += math.Pow(x[i]-math.Cos(0.5*math.Pi*x[0]), 2)
	}
	return s
}

func cmop1C1(x []float64) bool {
	return (0.51-cmop1G1(x))*(cmop1G1(x)-0.5) >= 0
}

func cmop1C2(x []float64) bool {
	return (0.51-cmop1G2(x))*(cmop1G2(x)-0.5) >= 0
}

func Cmop1(x []float64) Fitness {
	fitness := Fitness{}
	fitness.Objectives = make(map[string]float64)
	fitness.HardConstraints = make(map[string]bool)
	fitness.Objectives["f1"] = cmop1F1(x)
	fitness.Objectives["f2"] = cmop1F2(x)
	fitness.HardConstraints["c1"] = cmop1C1(x)
	fitness.HardConstraints["c2"] = cmop1C2(x)
	return fitness
}
