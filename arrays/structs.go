package arrays

import "math"

type Vector struct {
	Size  int
	inner []float64
}

func (v *Vector) Zeros() {
	v.inner = make([]float64, v.Size)
}

func (v *Vector) Fill(n float64) {
	v.inner = make([]float64, v.Size)
	for i := range v.inner {
		v.inner[i] = n
	}
}

func (v *Vector) Set(i int, val float64) {
	v.inner[i] = val
}

func (v *Vector) Mult(n float64) {
	for i := range v.inner {
		v.inner[i] *= n
	}
}

func (v *Vector) Length() float64 {
	var s float64
	for i := range v.inner {
		s += math.Pow(v.inner[i], 2)
	}
	return math.Sqrt(s)
}

func (v Vector) Sum() float64 {
	return Sum(v.inner)
}

func (v *Vector) Dist(u Vector) float64 {
	return EuclideanDistance(v.inner, u.inner)
}

func (v *Vector) Get(i int) float64 {
	return v.inner[i]
}
