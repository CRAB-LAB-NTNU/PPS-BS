package optimisers

import (
	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

/*Moead is the struct describing the MOEA/D algorithm.
 */
type Moead struct {
	archive, archiveCopy, population                []types.Individual
	cmop                                            types.Cmop
	chm                                             types.CHM
	T, WeightDistribution, populationSize           int
	DecisionVariables, Nr, generation, maxFuncEvals int
	fnEval, historyCounter                          int
	F, Cr, DistributionIndex, maxViolation          float64
	Weights                                         []arrays.Vector
	WeightNeigbourhood                              [][]int
	idealPoint                                      []float64
	binaryPairs                                     []int
}

func NewMoead(cmop types.Cmop, chm types.CHM, t, weightDistribution, decisionVariables, nr int, f, cr, distributionIndex float64, maxFuncEvals int) *Moead {

	moead := Moead{
		cmop:               cmop,
		chm:                chm,
		T:                  t,
		WeightDistribution: weightDistribution,
		DecisionVariables:  decisionVariables,
		Nr:                 nr,
		F:                  f,
		Cr:                 cr,
		DistributionIndex:  distributionIndex,
		maxFuncEvals:       maxFuncEvals,
	}

	moead.Initialise()
	return &moead
}

func (m Moead) Ideal() []float64 {
	return m.idealPoint
}

func (m Moead) MaxFuncEvals() int {
	return m.maxFuncEvals
}

func (m Moead) Archive() []types.Individual {
	return m.archive
}

func (m Moead) CHM() types.CHM {
	return m.chm
}

func (m Moead) MaxViolation() float64 {
	return m.maxViolation
}

func (m Moead) Population() []types.Individual {
	return m.population
}

func (m Moead) Generation() int {
	return m.generation
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
	m.Weights = arrays.UniformDistributedVectors(m.cmop.ObjectiveCount, m.WeightDistribution)

	m.populationSize = len(m.Weights)

	for i := range m.Weights {
		m.WeightNeigbourhood = append(m.WeightNeigbourhood, arrays.NearestNeighbour(m.Weights, i, m.T))
	}

	for i := 0; i < m.populationSize; i++ {
		ind := MoeadIndividual{Cmop: m.cmop}
		ind.Initialise()
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
		a[i] = ind.Fitness().TotalViolation()
	}
	return a
}

// FeasibleRatio returns the ratio of feasible feasible individuals in the population
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
Based on the stage parameter different evolutionary steps are taken
*/
func (m *Moead) Evolve(stage types.Stage) {

	//If the stage is binary search we evolve the population only using binary search
	if stage.Type() == types.BinarySearch {
		m.evolveBinary(stage)
		return
	}

	//If the stage is not push we care about constraints
	if stage.Type() != types.Push {
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
		if stage.Type() == types.Push {
			m.updatePopulation(hood, offspring, m.replaceIgnoringConstraints)
		} else {
			m.updatePopulation(hood, offspring, m.replaceWithConstraints)
		}
	}

	m.generation++
	m.archive = ndSelect(m.archive, m.population, m.populationSize)

}
