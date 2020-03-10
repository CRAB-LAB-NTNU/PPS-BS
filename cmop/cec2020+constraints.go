package cmop

import "math"

func cmop1Constraint(f1, f2, l, p float64) float64 {
	a := math.Pow(0.5*math.Pi-2*math.Abs(l-0.25*math.Pi), 3)
	b := math.Pow(1+p*math.Sin(6*a), 2)
	return b - math.Pow(f1, 2) - math.Pow(f2, 2)
}

func cmop2Constraint(f1, f2, l float64) float64 {
	a := math.Pow(f1/(1+0.15*l), 2)
	b := math.Pow(f2/(1+0.75*l), 2)
	return a + b - 1
}
