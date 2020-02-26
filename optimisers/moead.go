package optimisers

import (
	"math/rand"

	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
)

/*Moead is the struct describing the MOEA/D algorithm.
 */
type Moead struct {
	archive, ArchiveCopy, population                                     []types.Individual
	CMOP                                                                 types.CMOP
	WeightNeigbourhoodSize, WeightDistribution, populationSize           int
	DecisionSize, MaxChangeIndividuals, generation, GenerationMax        int
	fnEval, historyCounter                                               int
	DEDifferentialWeight, CrossoverRate, DistributionIndex, maxViolation float64
	Weights                                                              []arrays.Vector
	WeightNeigbourhood                                                   [][]int
	idealPoint                                                           []float64
	BinaryPairs                                                          []int
	BinaryMinDistance                                                    float64
	binaryCompleted                                                      bool
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
	m.historyCounter = -1
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
func (m *Moead) Evolve(stage types.Stage, doBinary bool, eps []float64) {

	if len(m.BinaryPairs) == 0 && doBinary && len(m.archive) > 0 {
		m.BinaryPairs = m.selectRandomPairs()
		m.ArchiveCopy = m.archiveCopy()
	}

	if len(m.BinaryPairs) > 0 {
		m.binarySearch()
		m.generation++
		return
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
		if constraintViolation(f) > m.maxViolation {
			m.maxViolation = constraintViolation(f)
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
	m.BinaryPairs = []int{}
}

func (m Moead) IsBinarySearch() bool {
	return len(m.BinaryPairs) != 0
}
