package testsuite2

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func inner1(x types.Genotype) float64 {
	var s float64
	for i := 1; i < len(x); i++ {
		t := x[i] - x[0]
		s += 0.5 * ((-0.9 * math.Pow(t, 2)) + math.Pow(math.Abs(t), 0.6))
	}
	return s
}
