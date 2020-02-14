package optimisers

import (
	"fmt"
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

	fmt.Println("Generating", m.populationSize, "Individuals")
	for i := 0; i < m.populationSize; i++ {
		ind := MoeadIndividual{D: m.DecisionSize}
		ind.InitialiseRandom(m.CMOP)
		m.population = append(m.population, &ind)
	}
	fmt.Println("Individuals:", len(m.population))
	fmt.Println("Calculating Ideal point")
	m.idealPoint = biooperators.CalculateIdealPoints(m.population)
	fmt.Println("Ideal Point:", m.idealPoint)
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
	feas, infeas := 0, 0
	for _, i := range m.population {
		if feasible(i.Fitness()) {
			feas++
		} else {
			infeas++
		}
	}
	if infeas == 0 {
		infeas = 1
	}
	return float64(feas) / float64(m.populationSize)
}

/*Evolve performs the genetic operator on all individuals in the population
 */
func (m *Moead) Evolve(stage types.Stage, eps []float64) {
	if m.generation%100 == 0 {
		fmt.Println("Evolving")
		fmt.Println("Generation:", m.generation, "Stage:", stage, "Archive Length:", len(m.archive), "eps:", eps[m.generation])
		fmt.Println("Ideal point:", m.idealPoint, "maxConstraintViolation:", m.maxViolation)
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
		for j, oType := range f.ObjectiveTypes {
			if oType == types.Minimisation && f.ObjectiveValues[j] < m.idealPoint[j] {
				m.idealPoint[j] = f.ObjectiveValues[j]
			} else if oType == types.Maximisation && f.ObjectiveValues[j] > m.idealPoint[j] {
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
