package metrics

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func InvertedGenerationalDistance(population []types.Individual, paretoFront [][]float64) (s float64) {
	for _, solution := range paretoFront {
		s += minDistance(population, solution)
	}
	return s / float64(len(paretoFront))
}

func distance(a, b []float64) (s float64) {
	for i := range a {
		s += math.Pow(a[i]-b[i], 2)
	}
	return math.Sqrt(s)
}

func minDistance(population []types.Individual, paretoPoint []float64) float64 {
	s := math.MaxFloat64
	for _, ind := range population {
		v := distance(ind.Fitness().ObjectiveValues, paretoPoint)
		if v < s {
			s = v
		}
	}
	return s
}
