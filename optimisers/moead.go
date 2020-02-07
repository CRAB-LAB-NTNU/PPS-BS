package optimisers

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
)

/*MoeadIndividual is a struct containing information about an individual in the population
of an evolutionary algorithm.
*/
type MoeadIndividual struct {
	D        int
	genotype types.Genotype
	fitness  types.Fitness
}

func (i *MoeadIndividual) SetGenotype(g []float64) {
	i.genotype = g
}

/*Genotype returns the genotype of the individual
 */
func (i MoeadIndividual) Genotype() types.Genotype {
	return i.genotype
}

/*Fitness returnes the fitness value of the individual if it has been calculated
 */
func (i MoeadIndividual) Fitness() types.Fitness {
	return i.fitness
}

/*UpdateFitness updates the fitness value of an individual
 */
func (i *MoeadIndividual) UpdateFitness(cmop types.CMOP) types.Fitness {
	i.fitness = cmop.Calculate(i.Genotype())
	return i.Fitness()
}

/*InitialiseRandom initialises the individuals genotype with random floats in the range [0,1]
 */
func (i *MoeadIndividual) InitialiseRandom(cmop types.CMOP) {
	i.genotype = make([]float64, i.D)
	for j := 0; j < i.D; j++ {
		i.genotype[j] = rand.Float64()
	}
	i.UpdateFitness(cmop)
}

func (ind *MoeadIndividual) Repair() {
	for i := 0; i < ind.D; i++ {
		if ind.genotype[i] > 1 {
			ind.genotype[i] = 1
		} else if ind.genotype[i] < 0 {
			ind.genotype[i] = 0
		}
	}
}

/*PolynomialMutation performs the mutation described in the
paper https://reader.elsevier.com/reader/sd/pii/S0045782599003898
PPS describes a mutation probability of 1/n where n => length of genotype.
We don't iterate but make a 1/n dice roll to check if we're mutating a single alele.
*/
func (i *MoeadIndividual) PolynomialMutation(m float64) {
	if m < 0 {
		panic("m needs to be non-negative")
	}

	if rand.Float64() > 1/float64(i.D) {
		return
	}

	u, r, gresk := rand.Float64(), rand.Intn(i.D), 0.0

	if u < 0.5 {
		gresk = math.Pow(2.0*u, 1.0/(m+1.0))
	} else {
		gresk = math.Pow(1.0-(2.0*(1.0-u)), 1.0/(m+1.0))
	}

	i.genotype[r] += gresk
}

/*Moead is the struct describing the MOEA/D algorithm.
 */
type Moead struct {
	Archive                                                              []types.Individual
	Population                                                           []types.Individual
	CMOP                                                                 types.CMOP
	WeightNeigbourhoodSize, WeightDistribution, populationSize           int
	DecisionSize, MaxChangeIndividuals                                   int
	DEDifferentialWeight, CrossoverRate, DistributionIndex, maxViolation float64
	Weights                                                              []arrays.Vector
	WeightNeigbourhood                                                   [][]int
	IdealPoint                                                           []float64
}

/*Initialise initialises the MOEA/D by calculating the weights, weight neighbourhood,
population and ideal point.
*/
func (m *Moead) Initialise() {
	fmt.Println("Initalising MOEA/D")
	fmt.Println("Creating weights")
	m.Weights = arrays.UniformDistributedVectors(m.CMOP.NumberOfObjectives, m.WeightDistribution)
	fmt.Println("Created", len(m.Weights))

	m.populationSize = len(m.Weights)

	fmt.Println("Calculating weight neighbourhood")
	for i := range m.Weights {
		m.WeightNeigbourhood = append(m.WeightNeigbourhood, arrays.NearestNeighbour(m.Weights, i, m.WeightNeigbourhoodSize))
	}
	//fmt.Println("Weight neighbourhood:", m.WeightNeigbourhood)

	fmt.Println("Generating", m.populationSize, "Individuals")
	for i := 0; i < m.populationSize; i++ {
		ind := MoeadIndividual{D: m.DecisionSize}
		ind.InitialiseRandom(m.CMOP)
		m.Population = append(m.Population, &ind)
	}
	//fmt.Println("Population:", m.Population)
	fmt.Println("Individuals:", len(m.Population))
	fmt.Println("Calculating Ideal point")
	m.IdealPoint = biooperators.CalculateIdealPoints(m.Population)
	fmt.Println("Ideal Point:", m.IdealPoint)
	m.maxViolation = -1
}

/*Crossover uses the traditional DE operator to generate a new individual
Our MOEA/D is implemented with a single offspring and parents picked from the
weight neighbourhood.
*/
func (m *Moead) Crossover(parents []types.Individual) []types.Individual {
	x, a, b := parents[0], parents[1], parents[2]
	child := MoeadIndividual{D: len(x.Genotype()), genotype: make([]float64, len(x.Genotype()))}
	for i := range child.Genotype() {
		if rand.Float64() < m.CrossoverRate {
			child.Genotype()[i] = x.Genotype()[i] + m.DEDifferentialWeight*(a.Genotype()[i]-b.Genotype()[i])
		} else {
			child.Genotype()[i] = x.Genotype()[i]
		}
	}

	child.PolynomialMutation(m.DistributionIndex)
	child.Repair()
	return []types.Individual{&child}
}

/*Evolve performs the genetic operator on all individuals in the population
 */
func (m *Moead) Evolve(stage types.Stage) {
	fmt.Println()
	fmt.Println("Evolving. Ideal point:", m.IdealPoint, "maxConstraintViolation:", m.maxViolation)
	for i := 0; i < m.populationSize; i++ {

		var hood []int

		if rand.Float64() < 0.9 {
			hood = make([]int, len(m.WeightNeigbourhood[i]))
			copy(hood, m.WeightNeigbourhood[i])
		} else {
			hood = make([]int, m.populationSize)
			for i := range hood {
				hood[i] = i
			}
		}

		x := rand.Intn(len(hood))
		if x == 0 {
			x = 1
		}
		y := rand.Intn(x)
		offSpring := m.Crossover([]types.Individual{m.Population[i], m.Population[hood[x]], m.Population[hood[y]]})

		offSpring[0].UpdateFitness(m.CMOP)

		// Update Ideal
		f := offSpring[0].Fitness()
		for j, oType := range f.ObjectiveTypes {
			if oType == types.Minimisation && f.ObjectiveValues[j] < m.IdealPoint[j] {
				m.IdealPoint[j] = f.ObjectiveValues[j]
			} else if oType == types.Maximisation && f.ObjectiveValues[j] > m.IdealPoint[j] {
				m.IdealPoint[j] = f.ObjectiveValues[j]
			}
		}

		// Update max violation
		if maximumConstraintViolation(f) > m.maxViolation {
			m.maxViolation = maximumConstraintViolation(f)
		}

		c := 0
		for c != m.MaxChangeIndividuals && len(hood) > 0 {
			j := rand.Intn(len(hood))
			replaced := false
			if stage == types.Push {
				parentScalar := tchebycheff(m.Population[hood[j]].Fitness().ObjectiveValues, m.IdealPoint, m.Weights[i])
				offSpringScalar := tchebycheff(f.ObjectiveValues, m.IdealPoint, m.Weights[i])
				if offSpringScalar <= parentScalar {
					replaced = true

					copyOffspring := MoeadIndividual{
						D:        len(offSpring[0].Genotype()),
						genotype: offSpring[0].Genotype(),
						fitness:  offSpring[0].Fitness(),
					}

					m.Population[hood[j]] = &copyOffspring

				} else {
					gene := make([]float64, len(m.Population[hood[j]].Genotype()))
					copy(gene, m.Population[hood[j]].Genotype())
					parentCopy := MoeadIndividual{
						D:        len(m.Population[hood[j]].Genotype()),
						genotype: gene,
						fitness:  m.Population[hood[j]].Fitness(),
					}
					m.Population[hood[j]] = &parentCopy
				}
			} else {
				// PULL
			}
			if replaced == true {
				c++
			}
			hood = arrays.Remove(hood, j)
		}
	}
	m.Archive = ndSelect(m.Archive, m.Population, m.populationSize)
}

func ndSelect(archive, population []types.Individual, n int) []types.Individual {
	fmt.Println("Selecting from archive of length:", len(archive), "and population of length:", len(population))

	union := biooperators.UnionPopulations(archive, population)
	fmt.Println("Created Union of length", len(union))

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

		for len(result)+len(q[i]) < n {
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
		sort.Slice(helper, func(i, j int) bool { return helper[i].Value < helper[j].Value })

		for j := 0; j < remaining; j++ {
			val := helper[j].Key
			result = append(result, q[i][val])
		}
	}
	return result
}

func maximumConstraintViolation(fitness types.Fitness) float64 {
	var s float64
	for _, cValue := range fitness.ConstraintValues {
		s += math.Abs(math.Min(cValue, 0))
	}
	return s
}

func feasible(fitness types.Fitness) bool {
	return maximumConstraintViolation(fitness) <= 0
}

func tchebycheff(objectiveValues, idealPoint []float64, weight arrays.Vector) float64 {
	var max float64 = math.SmallestNonzeroFloat64
	for i := range objectiveValues {
		// Original MOAD/D uses w(f-z)
		// PPS on the other hand uses 1/w * (f-z)
		// Using the formula from PPS.
		v := 1 / weight.Get(i) * (math.Abs(objectiveValues[i] - idealPoint[i]))
		if max < v {
			max = v
		}
	}
	return max
}

func Tchebycheff(objectiveValues, idealPoint []float64, weight arrays.Vector) float64 {
	return tchebycheff(objectiveValues, idealPoint, weight)
}
