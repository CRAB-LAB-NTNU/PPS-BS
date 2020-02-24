package testsuite2

import "github.com/CRAB-LAB-NTNU/PPS-BS/types"

func constraint1(x types.Genotype) float64 {
	g := inner1(x)
	if x[0] > 0.6 && g > 0.001 {
		return 10 * g
	}
	return 0
}

func constraint2(x types.Genotype) float64 {
	g := inner1(x)
	if x[0] < 0.2 || x[0] > 0.8 && g > 0.002 {
		return 10 * g
	}
	return 0
}

func constraint3(x types.Genotype) float64 {
	g := inner1(x)
	if x[0] < 0.2 || x[0] > 0.8 && g > 0.001 {
		return 10 * g
	}
	return 0
}
