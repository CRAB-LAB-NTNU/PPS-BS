package arrays

import (
	"math"
	"strconv"
)

func Includes(arr []int, val int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return true
		}
	}
	return false
}

func Remove(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func Sum(arr []float64) float64 {
	var s float64
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return s
}

func EuclideanDistance(a, b []float64) float64 {
	var s float64
	if len(a) != len(b) {
		panic("vectors not equal size")
	}
	for i := range a {
		s += math.Pow(a[i]-b[i], 2)
	}
	return math.Sqrt(s)
}

func floatArrayToString(a []float64) string {
	var s string
	for i := range a {
		s += strconv.FormatFloat(a[i], 'E', -1, 64)
		if i < len(a)-1 {
			s += ","
		}
	}
	return s
}

func IdentityMatrix(n int) [][]float64 {
	var matrix = make([][]float64, n)
	for i := 0; i < n; i++ {
		row := make([]float64, n)
		row[i] = 1
		matrix[i] = row
	}
	return matrix
}

func MiddleVector(a, b []float64) []float64 {
	middle := make([]float64, len(a))
	for i := range a {
		middle[i] = (a[i] + b[i]) / 2
	}
	return middle
}

func Normalise(a []float64) []float64 {
	s := Sum(a)
	var normalised = make([]float64, len(a))

	for i := range a {
		normalised[i] = a[i] / s
	}

	return normalised
}

//DistributedTriangleVectors creates a n*m matrix where n = objectives and m = population size.
//The algorithm creates new vertices in the n dimensional space until the row number is bigger than m.
//With this approach an even distribution is not always possible, but a good distribution can be reached.
func DistributedTriangleVectors(objectives, populationSize int) [][]float64 {

	vertices := IdentityMatrix(objectives)

	calculated := map[string]bool{}

	for i := range vertices {
		calculated[floatArrayToString(vertices[i])] = true
	}

	for len(vertices) < populationSize {
		k := len(vertices)
		for i := 0; i < k; i++ {
			for j := i; j < k; j++ {
				if i != j {
					vertex := MiddleVector(vertices[i], vertices[j])
					if calculated[floatArrayToString(vertex)] == true {
						continue
					} else {
						calculated[floatArrayToString(vertex)] = true
						vertices = append(vertices, vertex)
					}
				}
			}
		}
	}
	return vertices
}

//UniformDistributedVectors generates a set of uniformly distributed vectors
func UniformDistributedVectors(m, H int) []Vector {
	buf := make([]int, m)
	var result []Vector

	for true {
		vector := Vector{Size: m}
		vector.Zeros()
		for i := 0; i < m-1; i++ {
			vector.Set(i, float64(buf[i+1]-buf[i])/float64(H))
		}
		vector.Set(m-1, float64(H-buf[m-1])/float64(H))

		result = append(result, vector)

		var p int
		for p = m - 1; p != 0 && buf[p] == H; p-- {
		}
		if p == 0 {
			break
		}
		buf[p]++
		for p = p + 1; p != m; p++ {
			buf[p] = buf[p-1]
		}

	}
	return result[1 : len(result)-1]
}

//NearestNeighbour calculates the closest T vectors for vector i in the set arr.
//NOTE the implementation makes the assumption that the vector i is it's own closest neighbour.
//The MOEA/D framework also makes this assumption
func NearestNeighbour(arr []Vector, i, T int) []int {
	neighbourIndexes := make([]int, T)

	for x := range neighbourIndexes {
		neighbourIndexes[x] = -1
	}

	for x := 0; x < T; x++ {
		distance := math.MaxFloat64
		index := -1
		for j := range arr {
			dist := arr[i].Dist(arr[j])
			if dist < distance && !Includes(neighbourIndexes, j) {
				distance = dist
				index = j
			}
		}
		neighbourIndexes[x] = index
	}

	return neighbourIndexes
}

func Middle(a, b []float64) []float64 {
	middle := make([]float64, len(a))
	for i := range a {
		middle[i] = (a[i] + b[i]) / 2
	}
	return middle
}
