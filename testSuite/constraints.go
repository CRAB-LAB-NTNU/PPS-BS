package testSuite

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func Constraint1(x types.Genotype, g func(types.Genotype) float64) float64 {
	return (0.51 - g(x)) * (g(x) - 0.5)
}

func Constraint2(x types.Genotype) float64 {
	return math.Sin(20*math.Pi*x[0]) - 0.5
}

func Constraint3(x types.Genotype, p, q, a, b []float64, k int, f1, f2 float64) float64 {
	d := -0.25 * math.Pi
	z := math.Pow(((f1-p[k])*math.Cos(d)-(f2-q[k])*math.Sin(d))/a[k], 2)
	y := math.Pow(((f1-p[k])*math.Sin(d)+(f2-q[k])*math.Cos(d))/b[k], 2)
	return z + y - 0.1
}

func Constraint4(x types.Genotype, f1, f2, z float64) float64 {
	d := 0.25 * math.Pi
	a := f1*math.Sin(d) + f2*math.Cos(d)
	b := math.Sin(4 * math.Pi * (f1*math.Cos(d) - f2*math.Sin(d)))
	return a - b - z
}

func Constraint5(x types.Genotype, g, a, b float64) float64 {
	return (g - a) * (g - b)
}
