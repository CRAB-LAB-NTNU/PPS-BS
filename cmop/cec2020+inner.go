package cmop

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func g1(x types.Genotype) float64 {
	var s float64
	n := len(x)
	floatN := float64(n)
	for i := 1; i < n; i++ {
		floatI := float64(i)
		z := math.Pow(x[i], floatN-2)
		fraction := floatI / (2 * floatN)
		inner := math.Pow(z-0.5-fraction, 2)
		s += 1 - math.Exp(-10*inner)
	}
	return 1 + s
}

func g2(x types.Genotype) float64 {
	var s float64
	n := len(x)
	floatN := float64(n)

	for i := 1; i < n; i++ {
		floatI := float64(i)

		zExp := math.Pow(x[i]-(floatI/floatN), 2)

		z := 1 - math.Exp(-10*zExp)

		a := (0.1 / floatN) * math.Pow(z, 2)
		b := 1.5 * math.Cos(2*math.Pi*z)
		s += 1.5 + a - b
	}
	return 1 + s
}

func g3(x types.Genotype) float64 {
	var s float64
	for i := 1; i < len(x); i++ {
		v := x[i] + math.Pow(x[i-1]-0.5, 2) - 1
		s += math.Pow(v, 2)
	}
	return 1 + 2*s
}
