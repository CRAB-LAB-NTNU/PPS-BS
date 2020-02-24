package testSuite3

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func constraint1(x types.Genotype) float64 {
	return math.Sin(20 * math.Pi * x[0])
}
