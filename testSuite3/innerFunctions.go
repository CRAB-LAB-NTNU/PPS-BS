package testSuite3

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func inner1(x types.Genotype) float64 {
	var s float64
	for i := 1; i < len(x); i++ {
		t := x[i] - math.Sin(0.5*math.Pi*x[0])
		s += math.Pow(-0.9*t, 2) + math.Pow(math.Abs(t), 0.6)
	}
	return 2 * math.Sin(math.Pi*x[0]) * s
}

func inner2(x types.Genotype) float64 {
	var s float64
	for i := 1; i < len(x); i++ {
		t := x[i] - math.Sin(0.5*math.Pi*x[0])
		s += math.Abs(t) / (1 + math.Exp(5*math.Abs(t)))
	}
	return 10 * math.Sin(math.Pi*x[0]) * s
}

func inner3(x types.Genotype) float64 {
	var s float64
	for i := 1; i < len(x); i++ {
		t := x[i] - math.Sin(0.5*math.Pi*x[0])
		s += math.Abs(t) / (1 + math.Exp(5*math.Abs(t)))
	}
	return 10 * math.Sin((math.Pi*x[0])/2) * s
}

func inner4(x types.Genotype) float64 {
	var s float64
	for i := 1; i < len(x); i++ {
		t := x[i] - math.Sin(0.5*math.Pi*x[0])
		s += math.Abs(t) / (1 + math.Exp(5*math.Abs(t)))
	}
	return 1 + 10*math.Sin(math.Pi*x[0])*s
}
