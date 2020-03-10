package optimisers

import (
	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
)

/*Moead is the struct describing the MOEA/D algorithm.
 */
type Moead struct {
	archive, ArchiveCopy, population                []types.Individual
	CMOP                                            types.CMOP
	CHM                                             types.CHM
	T, WeightDistribution, populationSize           int
	DecisionVariables, Nr, generation, MaxFuncEvals int
	fnEval, historyCounter                          int
	F, Cr, DistributionIndex, maxViolation          float64
	Weights                                         []arrays.Vector
	WeightNeigbourhood                              [][]int
	idealPoint                                      []float64
	//Binary Search Attributes
	BinaryPairs                                           []int
	BinaryMinDistance, BinaryFeasibleSelectionProbability float64
	binaryCompleted                                       bool
}

func NewMoead(cmop types.CMOP, chm types.CHM, t, weightDistribution, decisionVariables, nr int, f, cr, distributionIndex float64, MaxFuncEvals int) *Moead {

	moead := Moead{
		CMOP:               cmop,
		CHM:                chm,
		T:                  t,
		WeightDistribution: weightDistribution,
		DecisionVariables:  decisionVariables,
		Nr:                 nr,
		F:                  f,
		Cr:                 cr,
		DistributionIndex:  distributionIndex,
		MaxFuncEvals:       MaxFuncEvals,
	}

	moead.Initialise()
	return &moead
}

func (m Moead) Ideal() []float64 {
	return m.idealPoint
}

func (m Moead) GetMaxFuncEvals() int {
	return m.MaxFuncEvals
}

func (m Moead) Archive() []types.Individual {
	return m.archive
}

func (m Moead) GetCHM() types.CHM {
	return m.CHM
}

/*
func (m Moead) MaxViolation() float64 {
	return m.maxViolation
}*/

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
		m.WeightNeigbourhood = append(m.WeightNeigbourhood, arrays.NearestNeighbour(m.Weights, i, m.T))
	}

	for i := 0; i < m.populationSize; i++ {
		ind := MoeadIndividual{D: m.DecisionVariables}
		ind.InitialiseRandom(m.CMOP)
		m.population = append(m.population, &ind)
	}
	m.idealPoint = biooperators.CalculateIdealPoints(m.population)
	m.maxViolation = -1
	m.historyCounter = -1
}

// FunctionEvaluations returns the current number of function evaluations performed
func (m Moead) FunctionEvaluations() int {
	return m.fnEval
}

//ConstraintViolation returns the total constraint violation of the population
func (m Moead) ConstraintViolation() []float64 {
	a := make([]float64, m.populationSize)
	for i, ind := range m.population {
		a[i] = constraintViolation(ind.Fitness())
	}
	return a
}

// FeasibleRatio returns the ratio of feasible feasible individuals in the population
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
Based on the stage parameter different evolutionary steps are taken
*/
func (m *Moead) Evolve(stage types.StageType, threshold float64) {

	//If the stage is binary search we evolve the population only using binary search
	if stage == types.BinarySearch {
		m.evolveBinary()
		return
	}

	//If the stage is not push we care about constraints
	if stage != types.Push {
		m.updateCHM()
	}

	for i := 0; i < m.populationSize; i++ {

		hood := m.selectHood(0.9, i)
		p := m.population[i]
		x, y := m.selectIndividualsForCrossover(hood)
		offspring := m.crossover(p, x, y)
		m.fnEval++

		m.updateIdealPoint(offspring)

		m.updateMaxConstraintViolation(offspring)

		//Would have preferred a more modular appraoch without the check
		if stage == types.Push {
			m.updatePopulation(hood, offspring, m.replaceIgnoringConstraints)
		} else {
			m.updatePopulation(hood, offspring, m.replaceWithConstraints)
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

func (m Moead) BinaryDone() bool {
	return m.binaryCompleted
}
