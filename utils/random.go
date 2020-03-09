package utils

import "math/rand"

//RandomFloat64Range returns a random number within minmax range.
func RandomFloat64Range(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
