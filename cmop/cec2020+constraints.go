package cmop

import "math"

func cmop2Constraint(f1, f2, l float64) float64 {
	a := math.Pow(f1/(1+0.15*l), 2)
	b := math.Pow(f2/(1+0.75*l), 2)
	return a + b - 1
}
