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

// Nå skal jeg ut på eventyr
// For jeg har hengt med Lars
// Og smitte noen rare dyr
// For Lars har gitt meg SARS
// Reiser over land og strand
// Ut på evig jakt
// Gi corona, se det går an
// Å bruke dennes makt
// LARS HAR SARS
