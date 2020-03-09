package testSuite

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func Inner1(x types.Genotype) float64 {
	var s float64
	for i := 2; i < len(x); i += 2 {
		s += math.Pow(x[i]-math.Sin(0.5*math.Pi*x[0]), 2)
	}
	return s
}

func Inner2(x types.Genotype) float64 {
	var s float64
	for i := 1; i < len(x); i += 2 {
		s += math.Pow(x[i]-math.Cos(0.5*math.Pi*x[0]), 2)
	}
	return s
}

func Inner3(x types.Genotype) float64 {
	var s float64
	for i := 2; i < len(x); i += 2 {
		k := 0.5 * float64(i) / 30
		s += math.Pow(x[i]-math.Sin(k*math.Pi*x[0]), 2)
	}
	return s
}

func Inner4(x types.Genotype) float64 {
	var s float64
	for i := 1; i < len(x); i += 2 {
		k := 0.5 * float64(i) / 30
		s += math.Pow(x[i]-math.Cos(k*math.Pi*x[0]), 2)
	}
	return s
}

func Inner5(x types.Genotype) float64 {
	var s float64
	for i := 2; i < len(x); i++ {
		s += 10 * math.Pow(x[i]-0.5, 2)
	}
	return s
}
