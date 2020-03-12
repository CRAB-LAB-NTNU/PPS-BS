package optimisers

import (
	"fmt"
	"math/rand"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

/*Moead is the struct describing the MOEA/D algorithm.
 */
type Moead struct {
	archive, ArchiveCopy, population                                     []types.Individual
	Cmop                                                                 types.CMOP
	WeightNeigbourhoodSize, WeightDistribution, populationSize           int
	MaxChangeIndividuals, generation, EvaluationsMax                     int
	fnEval, historyCounter                                               int
	DEDifferentialWeight, CrossoverRate, DistributionIndex, maxViolation float64
	BoundaryMinDistance, BoundaryFeasibleSelectionProbability            float64
	WeightNeigbourhood                                                   [][]int
	idealPoint                                                           []float64
	BoundaryPairs                                                        []int
	Weights                                                              []arrays.Vector
	binaryCompleted                                                      bool
}

func (m Moead) Ideal() []float64 {
	return m.idealPoint
}

func (m Moead) MaxEvaluations() int {
	return m.EvaluationsMax
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

}

/*Initialise initialises the MOEA/D by calculating the weights, weight neighbourhood,
population and ideal point.
*/
func (m *Moead) Initialise(cmop types.CMOP) {
	m.Cmop = cmop
	m.archive = make([]types.Individual, 0)
	m.population = make([]types.Individual, 0)
	m.WeightNeigbourhood = make([][]int, 0)
	m.idealPoint = make([]float64, 0)
	m.generation = 0
	m.fnEval = 0

	m.binaryCompleted = false

	m.Weights = arrays.UniformDistributedVectors(m.Cmop.ObjectiveCount, m.WeightDistribution)

	m.populationSize = len(m.Weights)

	for i := range m.Weights {
		m.WeightNeigbourhood = append(m.WeightNeigbourhood, arrays.NearestNeighbour(m.Weights, i, m.WeightNeigbourhoodSize))
	}

	for i := 0; i < m.populationSize; i++ {
		ind := MoeadIndividual{Cmop: m.Cmop}
		ind.Initialise()
		m.population = append(m.population, &ind)
	}
	m.idealPoint = biooperators.CalculateIdealPoints(m.population)
	m.maxViolation = -1
	m.historyCounter = -1
}

func (m Moead) FunctionEvaluations() int {
	return m.fnEval
}
func (m Moead) ConstraintViolation() []float64 {
	a := make([]float64, m.populationSize)
	for i, ind := range m.population {
		a[i] = ind.Fitness().ConstraintViolation()
	}
	return a
}

func (m Moead) FeasibleRatio() float64 {
	feas := 0
	for _, i := range m.population {
		if i.Fitness().Feasible() {
			feas++
		}
	}

	return float64(feas) / float64(m.populationSize)
}

/*Evolve performs the genetic operator on all individuals in the population
 */
func (m *Moead) Evolve(stage types.Stage, eps []float64) {
	if stage == types.BorderSearch {
		if len(m.archive) == 0 {
			fmt.Println("Skipping binary")
			m.binaryCompleted = true
		} else {
			if len(m.BoundaryPairs) == 0 {
				m.BoundaryPairs = m.selectRandomPairs()
				m.ArchiveCopy = m.archiveCopy()
			}
			m.boundarySearch()
			m.generation++
			return
		}
	}

	for i := 0; i < m.populationSize; i++ {
		hood := m.selectHood(0.9, i)
		x := rand.Intn(len(hood))
		y := x
		for y == x {
			y = rand.Intn(len(hood))
		}
		offSpring := m.crossover([]types.Individual{m.population[i], m.population[hood[x]], m.population[hood[y]]})[0]

		m.fnEval++
		// Update Ideal
		f := offSpring.Fitness()
		for j, val := range f.ObjectiveValues {
			if val < m.idealPoint[j] {
				m.idealPoint[j] = f.ObjectiveValues[j]
			}
		}
		// Update max violation
		if f.ConstraintViolation() > m.maxViolation {
			m.maxViolation = f.ConstraintViolation()
		}
		c := 0
		for c < m.MaxChangeIndividuals && len(hood) > 0 {
			j := rand.Intn(len(hood))
			replaced := false
			if stage == types.Push {
				replaced = m.pushProblems(hood[j], offSpring)
			} else {
				replaced = m.pullProblems(hood[j], offSpring, eps[m.generation])
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

func (m *Moead) ResetBinary() {
	m.ArchiveCopy = []types.Individual{}
	m.BoundaryPairs = []int{}
}

func (m Moead) IsBinarySearch() bool {
	return len(m.BoundaryPairs) != 0
}

func (m Moead) BinaryDone() bool {
	return m.binaryCompleted
}
