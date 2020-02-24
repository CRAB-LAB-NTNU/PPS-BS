package optimisers

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
)

/*Moead is the struct describing the MOEA/D algorithm.
 */
type Moead struct {
	archive                                                              []types.Individual
	population                                                           []types.Individual
	CMOP                                                                 types.CMOP
	WeightNeigbourhoodSize, WeightDistribution, populationSize           int
	DecisionSize, MaxChangeIndividuals, generation, GenerationMax        int
	fnEval                                                               int
	DEDifferentialWeight, CrossoverRate, DistributionIndex, maxViolation float64
	Weights                                                              []arrays.Vector
	WeightNeigbourhood                                                   [][]int
	idealPoint                                                           []float64
	binaryPairs                                                          []int
}

func (m Moead) Ideal() []float64 {
	return m.idealPoint
}

func (m Moead) MaxGeneration() int {
	return m.GenerationMax
}

func (m Moead) Archive() []types.Individual {
	return m.archive
}

func (m Moead) MaxViolation() float64 {
	return m.maxViolation
}
func (m Moead) Population() []types.Individual {
	return m.population
}

func (m *Moead) Reset() {
	m.archive = []types.Individual{}
	m.population = []types.Individual{}
	m.generation = 0
	m.fnEval = 0
	m.Weights = []arrays.Vector{}
	m.WeightNeigbourhood = [][]int{}
	m.idealPoint = []float64{}
}

/*Initialise initialises the MOEA/D by calculating the weights, weight neighbourhood,
population and ideal point.
*/
func (m *Moead) Initialise() {
	m.Weights = arrays.UniformDistributedVectors(m.CMOP.NumberOfObjectives, m.WeightDistribution)

	m.populationSize = len(m.Weights)

	for i := range m.Weights {
		m.WeightNeigbourhood = append(m.WeightNeigbourhood, arrays.NearestNeighbour(m.Weights, i, m.WeightNeigbourhoodSize))
	}

	for i := 0; i < m.populationSize; i++ {
		ind := MoeadIndividual{D: m.DecisionSize}
		ind.InitialiseRandom(m.CMOP)
		m.population = append(m.population, &ind)
	}
	m.idealPoint = biooperators.CalculateIdealPoints(m.population)
	m.maxViolation = -1
}

/*Crossover uses the traditional DE operator to generate a new individual
Our MOEA/D is implemented with a single offspring and parents picked from the
weight neighbourhood.
*/
func (m *Moead) Crossover(parents []types.Individual) []types.Individual {

	x, a, b := parents[0], parents[1], parents[2]

	child := MoeadIndividual{
		D:        m.DecisionSize,
		genotype: make([]float64, m.DecisionSize),
	}

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
func (m Moead) FunctionEvaluations() int {
	return m.fnEval
}
func (m Moead) ConstraintViolation() []float64 {
	a := make([]float64, m.populationSize)
	for i, ind := range m.population {
		a[i] = constraintViolation(ind.Fitness())
	}
	return a
}

func (m Moead) FeasibleRatio() float64 {
	feas := 0
	for _, i := range m.population {
		if feasible(i.Fitness()) {
			feas++
		}
	}

	return float64(feas) / float64(m.populationSize)
}

/*Evolve performs the genetic operator on all individuals in the population
 */
func (m *Moead) Evolve(stage types.Stage, eps []float64) {

	if m.generation%10 == 0 {
		fmt.Println(m.generation /*, arrays.Sum(m.ConstraintViolation())*/)
	}

	for i := 0; i < m.populationSize; i++ {
		hood := m.selectHood(0.9, i)
		x := rand.Intn(len(hood))
		y := x
		for y == x {
			y = rand.Intn(len(hood))
		}
		offSpring := m.Crossover([]types.Individual{m.population[i], m.population[hood[x]], m.population[hood[y]]})[0]
		offSpring.UpdateFitness(m.CMOP)
		m.fnEval++
		// Update Ideal
		f := offSpring.Fitness()
		for j, val := range f.ObjectiveValues {
			if val < m.idealPoint[j] {
				m.idealPoint[j] = f.ObjectiveValues[j]
			}
		}
		// Update max violation
		if constraintViolation(f) > m.maxViolation {
			m.maxViolation = constraintViolation(f)
		}
		c := 0
		for c < m.MaxChangeIndividuals && len(hood) > 0 {
			j := rand.Intn(len(hood))
			replaced := false
			if stage == types.Push {
				replaced = m.PushProblems(hood[j], offSpring)
			} else {
				replaced = m.PullProblems(hood[j], offSpring, eps[m.generation])
			}
			if replaced == true {
				c++
			}
			hood = arrays.Remove(hood, j)
		}
	}

	m.generation++
	m.archive = ndSelect(m.archive, m.population, m.populationSize)
}

func (m *Moead) PushProblems(j int, y types.Individual) bool {
	xF := m.population[j].Fitness()
	yF := y.Fitness()

	xS := tchebycheff(xF.ObjectiveValues, m.idealPoint, m.Weights[j])
	yS := tchebycheff(yF.ObjectiveValues, m.idealPoint, m.Weights[j])

	if yS <= xS {
		m.population[j] = y
		return true
	}
	return false
}

func (m *Moead) PullProblems(j int, y types.Individual, eps float64) bool {
	xF := m.population[j].Fitness()
	yF := y.Fitness()
	xCV := constraintViolation(xF)
	yCV := constraintViolation(yF)
	xS := tchebycheff(xF.ObjectiveValues, m.idealPoint, m.Weights[j])
	yS := tchebycheff(yF.ObjectiveValues, m.idealPoint, m.Weights[j])

	if yCV <= eps && xCV <= eps {
		if yS <= xS {
			m.population[j] = y
			return true
		}
	} else if yCV == xCV {
		if yS <= xS {
			m.population[j] = y
			return true
		}
	} else if yCV < xCV {
		m.population[j] = y
		return true
	}

	return false
}

func (m Moead) selectHood(pr float64, i int) []int {
	if rand.Float64() < pr {
		hood := make([]int, m.WeightNeigbourhoodSize)
		copy(hood, m.WeightNeigbourhood[i])
		return hood
	}
	hood := make([]int, m.populationSize)
	for i := range hood {
		hood[i] = i
	}
	return hood
}

func (m *Moead) binarySearch(indices []int, surrogate []types.Individual, eps float64) {
	if len(indices) != len(surrogate) {
		log.Fatal("INDICES AND SURROGATE NOT EQUAL LENGTH")
	}
	for i, p := range m.population {
		pair := surrogate[indices[i]]
		middlePoint := arrays.Middle(p.Genotype(), pair.Genotype())
		ind := MoeadIndividual{D: len(p.Genotype())}
		ind.SetGenotype(middlePoint)
		ind.Repair()
		ind.UpdateFitness(m.CMOP)
		m.fnEval++
		m.PullProblems(i, &ind, eps)
	}
}

func (m Moead) selectRandomPairs() []int {
	indices := make([]int, len(m.population))
	for i := range m.population {
		indices[i] = rand.Intn(len(m.archive))
	}
	return indices
}

func (m Moead) selectClosestPairs() []int {
	indices := make([]int, len(m.population))
	for i, p := range m.population {
		smallest := math.MaxFloat64
		flag := -1
		for j, a := range m.archive {
			dist := arrays.EuclideanDistance(p.Fitness().ObjectiveValues, a.Fitness().ObjectiveValues)
			if dist < smallest {
				flag = j
				smallest = dist
			}
		}
		indices[i] = flag
	}
	return indices
}

func (m Moead) selectFurthestPairs() []int {
	indices := make([]int, len(m.population))
	for i, p := range m.population {
		smallest := math.SmallestNonzeroFloat64
		for j, a := range m.population {
			if arrays.Includes(indices, j) {
				continue
			}
			dist := arrays.EuclideanDistance(p.Fitness().ObjectiveValues, a.Fitness().ObjectiveValues)
			if dist > smallest {
				indices[i] = j
				smallest = dist
			}
		}
	}
	return indices
}

func (m Moead) PopulationCentroid() []float64 {
	c := make([]float64, len(m.Population()[0].Genotype()))
	for _, ind := range m.Population() {
		for j := range ind.Genotype() {
			c[j] += ind.Genotype()[j]
		}
	}
	for j := range c {
		c[j] = c[j] / float64(m.populationSize)
	}
	return c
}

func (m Moead) OppositePopulation() []types.Individual {
	var pop []types.Individual
	centroid := m.PopulationCentroid()
	for _, ind := range m.Population() {
		d := len(ind.Genotype())
		oppositeInd := MoeadIndividual{D: d, genotype: make([]float64, d)}
		for j := range oppositeInd.genotype {
			oppositeInd.genotype[j] = 2*centroid[j] - ind.Genotype()[j]
		}
		oppositeInd.Repair()
		oppositeInd.UpdateFitness(m.CMOP)
		pop = append(pop, &oppositeInd)
	}
	return pop
}
