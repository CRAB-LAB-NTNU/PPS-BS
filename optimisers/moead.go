package optimisers

import (
	"fmt"
	"math/rand"

	"github.com/CRAB-LAB-NTNU/PPS-BS/r2s"

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
	BoundaryPairs                                                        []int
	BoundaryMinDistance, BoundaryFeasibleSelectionProbability            float64
	binaryCompleted                                                      bool
	R2s                                                                  r2s.R2S
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
	m.binaryCompleted = false
}

/*Initialise initialises the MOEA/D by calculating the weights, weight neighbourhood,
population and ideal point.
*/
func (m *Moead) Initialise() {
	m.Weights = arrays.UniformDistributedVectors(m.CMOP.NumberOfObjectives(), m.WeightDistribution)

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

	m.R2s.Initialize(m.MaxGeneration(), 3, 1000, m.FeasibleRatio(), m.Population())
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
		x, y := m.selectIndividualsForCrossover(hood)
		offSpring := m.crossover([]types.Individual{m.population[i], x, y})[0]

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

func (m *Moead) EvolveR2s() {

	offSpring := make([]types.Individual, m.populationSize)
	hoods := make([][]int, m.populationSize)
	for i := 0; i < m.populationSize; i++ {
		hood := m.selectHood(0.9, i)
		hoods[i] = hood
		x, y := m.selectIndividualsForCrossover(hood)
		offSpring[i] = m.crossover([]types.Individual{m.population[i], x, y})[0]

		m.fnEval++

		f := offSpring[i].Fitness()
		for j, val := range f.ObjectiveValues {
			if val < m.idealPoint[j] {
				m.idealPoint[j] = f.ObjectiveValues[j]
			}
		}

	}

	rankedPopulation := biooperators.FastNonDominatedSort(m.population)
	randomIndex := rand.Intn(len(rankedPopulation[0]))
	randomBest := rankedPopulation[0][randomIndex]

	m.R2s.ACD(m.generation, m.fnEval, randomBest.Fitness())

	for _, activeConstraint := range m.R2s.ActiveConstraints {
		if activeConstraint {
			m.R2s.UpdateDeltaIn(m.generation, m.fnEval)
			m.R2s.UpdateDeltaOut(m.generation, m.fnEval)
			break
		} else {
			//TODO remove boundary??
		}
	}

	newPop := make([]types.Individual, m.populationSize)

	for i := range m.population {

		hood := hoods[i]

		oF := offSpring[i].Fitness()
		pF := m.population[i].Fitness()
		oCV := m.R2s.ConstraintViolation(m.generation, oF)
		pCV := m.R2s.ConstraintViolation(m.generation, pF)

		for len(hood) > 0 {
			j := rand.Intn(len(hood))
			oS := tchebycheff(oF.ObjectiveValues, m.idealPoint, m.Weights[j])
			pS := tchebycheff(pF.ObjectiveValues, m.idealPoint, m.Weights[j])

			if oCV <= 0 && pCV <= 0 {
				if oS <= pS {
					newPop[i] = offSpring[i]
					break
				} else {
					newPop[i] = m.population[i]
				}
			} else if oCV == pCV {
				if oS <= pS {
					newPop[i] = offSpring[i]
					break
				} else {
					newPop[i] = m.population[i]
				}
			} else if oCV < pCV {
				newPop[i] = offSpring[i]
				break
			} else {
				newPop[i] = m.population[i]
			}

			hood = arrays.Remove(hood, j)
		}

	}
	m.population = newPop
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
