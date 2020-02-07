package biooperators

import (
	"fmt"
	"math"
	"sort"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

/*CalculateIdealPoints calculates the ideal point in a population,
IE. the point in the search space by picking the best function value for all objective functions.
*/
func CalculateIdealPoints(population []types.Individual) []float64 {
	oc := population[0].Fitness().ObjectiveCount
	point := make([]float64, oc)
	for i := 0; i < oc; i++ {
		if population[0].Fitness().ObjectiveTypes[i] == types.Minimisation {
			point[i] = math.MaxFloat64
		} else {
			point[i] = math.SmallestNonzeroFloat64
		}
	}
	for _, ind := range population {
		fitness := ind.Fitness()
		for j := 0; j < fitness.ObjectiveCount; j++ {
			if fitness.ObjectiveTypes[j] == types.Minimisation && fitness.ObjectiveValues[j] < point[j] {
				point[j] = fitness.ObjectiveValues[j]
			} else if fitness.ObjectiveTypes[j] == types.Maximisation && fitness.ObjectiveValues[j] > point[j] {
				point[j] = fitness.ObjectiveValues[j]
			}
		}
	}
	return point
}

func FastNonDominatedSort(population []types.Individual) [][]types.Individual {

	fronts := make([][]types.Individual, len(population))
	sets := make([][]types.Individual, len(population))
	n := make([]int, len(population))
	indexLookup := make(map[types.Individual]int)

	for p := range population {
		indexLookup[population[p]] = p
		for q := range population {
			if Dominates(population[p], population[q]) {
				sets[p] = append(sets[p], population[q])
			} else if Dominates(population[q], population[p]) {
				n[p]++
			}
		}
		if n[p] == 0 {
			fronts[0] = append(fronts[0], population[p])
		}
	}
	i := 0
	for len(fronts[i]) > 0 {
		var H []types.Individual
		for p := range fronts[i] {
			p = indexLookup[fronts[i][p]]
			for q := range sets[p] {
				q = indexLookup[sets[p][q]]
				n[q]--
				if n[q] == 0 {
					H = append(H, population[q])
				}
			}
		}
		i++
		fronts[i] = H
	}
	for j := len(fronts) - 1; j >= 0; j-- {
		if len(fronts[j]) > 0 {
			break
		}
		fronts = fronts[:j]
	}
	return fronts
}

func Dominates(p, q types.Individual) bool {
	qf, pf := q.Fitness(), p.Fitness()
	for i, oType := range qf.ObjectiveTypes {
		if oType == types.Minimisation {
			if qf.ObjectiveValues[i] < pf.ObjectiveValues[i] {
				return false
			}
		} else {
			if qf.ObjectiveValues[i] > pf.ObjectiveValues[i] {
				return false
			}
		}
	}
	return true
}

func CrowdingDistance(population []types.Individual) map[int]float64 {

	lookup := make(map[types.Individual]int)

	for i, p := range population {
		lookup[p] = i
	}

	distances := make(map[int]float64)
	l := len(population) - 1

	for m := range population[0].Fitness().ObjectiveValues {
		sorted := SortByValue(population, m)
		distances[lookup[sorted[0]]], distances[lookup[sorted[l]]] = math.MaxFloat64, math.MaxFloat64
		for i := 1; i < l-1; i++ {
			distances[lookup[sorted[i]]] += sorted[i+1].Fitness().ObjectiveValues[m] - sorted[i-1].Fitness().ObjectiveValues[m]
		}
	}
	return distances
}

func SortByValue(population []types.Individual, index int) []types.Individual {
	sorted := make([]types.Individual, len(population))
	copy(sorted, population)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Fitness().ObjectiveValues[index] < sorted[j].Fitness().ObjectiveValues[index]
	})
	return sorted
}

func PrintIndividualToGeogebraPoint(ind types.Individual) {
	fmt.Print("(")
	for i, m := range ind.Fitness().ObjectiveValues {
		fmt.Print(m)
		if i != len(ind.Fitness().ObjectiveValues)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println(")")
}

func UnionPopulations(a, b []types.Individual) []types.Individual {
	check := make(map[types.Individual]bool)
	var union []types.Individual
	counter := 0
	for _, ind := range a {
		if _, ok := check[ind]; ok {
			fmt.Println("Union exists")
			counter++
		}
		check[ind] = true
		union = append(union, ind)
	}
	for _, ind := range b {
		if _, ok := check[ind]; !ok {
			union = append(union, ind)
		} else {
			fmt.Println("Union duplicate")
		}
	}
	return union
}
