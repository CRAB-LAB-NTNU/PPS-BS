package optimisers

import (
	"math"
	"sort"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func ndSelect(archive, population []types.Individual, n int) []types.Individual {
	//fmt.Println("Selecting from archive of length:", len(archive), "and population of length:", len(population))

	union := biooperators.UnionPopulations(archive, population)
	//fmt.Println("Created Union of length", len(union))

	var feasibleSet []types.Individual
	feasibleCount := 0
	var result []types.Individual

	for i, ind := range union {
		if feasible(ind.Fitness()) {
			feasibleCount++
			feasibleSet = append(feasibleSet, union[i])
		}
	}
	if feasibleCount <= n {
		result = feasibleSet
	} else {
		q := biooperators.FastNonDominatedSort(feasibleSet)
		i := 0

		for len(result)+len(q[i]) < n && i < len(q)-1 {
			result = append(result, q[i]...)
			i++
		}
		remaining := n - len(result)
		distances := biooperators.CrowdingDistance(q[i])
		type sa struct {
			Key   int
			Value float64
		}
		var helper []sa
		for k, v := range distances {
			helper = append(helper, sa{k, v})
		}
		sort.Slice(helper, func(i, j int) bool { return helper[i].Value > helper[j].Value })

		for j := 0; j < remaining && j < len(helper); j++ {
			val := helper[j].Key
			result = append(result, q[i][val])
		}
	}
	return result
}

func constraintViolation(fitness types.Fitness) float64 {
	var s float64
	for _, cValue := range fitness.ConstraintValues {
		if cValue < 0 {
			s += math.Abs(cValue)
		}
	}
	return s
}

func feasible(fitness types.Fitness) bool {
	return constraintViolation(fitness) <= 0
}

func tchebycheff(objectiveValues, idealPoint []float64, weight arrays.Vector) float64 {
	var max float64 = math.SmallestNonzeroFloat64
	for i := range objectiveValues {
		// Original MOAD/D uses w(f-z)
		// PPS on the other hand uses 1/w * (f-z)
		// Using the formula from PPS.
		v := math.Abs(objectiveValues[i]-idealPoint[i]) / weight.Get(i)
		if max < v {
			max = v
		}
	}
	return max
}
