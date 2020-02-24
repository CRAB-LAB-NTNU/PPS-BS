package types

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
)

type Results struct {
	values []float64
}

func (r *Results) Add(v float64) {
	r.values = append(r.values, v)
}

func (r Results) Mean() float64 {
	return arrays.Sum(r.values) / float64(len(r.values))
}

func (r Results) Variance() float64 {
	mean := r.Mean()
	var s float64
	for i := range r.values {
		s += math.Pow(r.values[i]-mean, 2)
	}
	return s / float64(len(r.values)-1)
}

func (r Results) StandardDeviation() float64 {
	return math.Sqrt(r.Variance())
}

func (r Results) Values() []float64 {
	return r.values
}
