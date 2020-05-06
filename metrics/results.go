package metrics

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

type metricData []float64

func (metric metricData) size() float64 {
	return float64(len(metric))
}
func (metric metricData) Last() float64 {
	if len(metric) == 0 {
		return math.NaN()
	}
	return metric[len(metric)-1]
}
func (metric metricData) Mean() float64 {
	return arrays.Sum(metric) / metric.size()
}

func (metric metricData) Variance() float64 {
	mean := metric.Mean()
	var s float64
	for i := range metric {
		s += math.Pow(metric[i]-mean, 2)
	}
	return s / (metric.size() - 1)
}

func (metric metricData) StandardDeviation() float64 {
	return math.Sqrt(metric.Variance())
}

type Results struct {
	ParetoFront          [][]float64
	archives             [][]types.Individual
	IGD, HV              metricData
	HyperVolumeReference func() []float64
}

func (r *Results) Add(archive []types.Individual) {
	r.archives = append(r.archives, archive)
	if len(archive) == 0 {
		return
	}
	r.IGD = append(r.IGD, InvertedGenerationalDistance(archive, r.ParetoFront))
	r.HV = append(r.HV, HyperVolume(archive, r.HyperVolumeReference()))
}

func (r Results) FeasibilityRate() float64 {
	return r.IGD.size() / float64(len(r.archives))
}
