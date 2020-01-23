package testSuite

import "math"

func cmop2F1(x []float64) float64 {
	return x[0] + cmop1G1(x)
}

func cmop2F2(x []float64) float64 {
	return 1 - math.Sqrt(x[0]) + cmop1G2(x)
}

func Cmop2(x []float64) Fitness {
	fitness := Fitness{}
	fitness.Objectives = make(map[string]float64)
	fitness.HardConstraints = make(map[string]bool)
	fitness.Objectives["f1"] = cmop2F1(x)
	fitness.Objectives["f2"] = cmop2F2(x)
	fitness.HardConstraints["c1"] = cmop1C1(x)
	fitness.HardConstraints["c2"] = cmop1C2(x)
	return fitness
}
