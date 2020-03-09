package CEC2020

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func CEC_CMOP1(x types.Genotype) types.Fitness {
	fitness := types.Fitness{
		ObjectiveCount:  2,
		ConstraintCount: 3,
	}

	g := g1(x)
	f1 := x[0] * g
	f2 := g * math.Sqrt(1-math.Pow(f1/g, 2))

	l := math.Atan(f2 / f1)
	c1 := math.Pow(f1, 2) + math.Pow(f2, 2) - math.Pow(1.7-0.2*math.Sin(2*l), 2)
	c2 := c(f1, f2, l, 0.5)
	c3 := c(f1, f2, l, 0.45)

	fitness.ObjectiveValues = []float64{
		f1,
		f2,
	}
	fitness.ConstraintValues = []float64{
		c1,
		c2,
		c3,
	}
	fitness.ConstraintTypes = []types.ConstraintType{
		types.EqualsOrLessThanZero,
		types.EqualsOrLessThanZero,
		types.EqualsOrLessThanZero,
	}
	return fitness
}

func c(f1, f2, l, p float64) float64 {
	a := math.Pow(0.5*math.Pi-2*math.Abs(l-0.25*math.Pi), 3)
	b := math.Pow(1+p*math.Sin(6*a), 2)
	return b - math.Pow(f1, 2) - math.Pow(f2, 2)
}
