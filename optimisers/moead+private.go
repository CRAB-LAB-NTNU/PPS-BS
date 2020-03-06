package optimisers

import (
	"math/rand"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

/*crossover uses the traditional DE operator to generate a new individual
Our MOEA/D is implemented with a single offspring and parents picked from the
weight neighbourhood.
*/
func (m *Moead) crossover(parents []types.Individual) []types.Individual {

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
	child.UpdateFitness(m.CMOP)
	return []types.Individual{&child}
}

func (m *Moead) pushProblems(j int, y types.Individual) bool {
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

func (m *Moead) pullProblems(j int, y types.Individual, eps float64) bool {
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

func (m *Moead) selectIndividualsForCrossover(hood []int) (types.Individual, types.Individual) {
	x := rand.Intn(len(hood))
	y := x
	for y == x {
		y = rand.Intn(len(hood))
	}
	return m.population[hood[x]], m.population[hood[y]]

}

func (m *Moead) boundarySearch() {
	missCounter := 0
	for i, p := range m.population {
		m.fnEval++
		j := m.BoundaryPairs[i]
		if j == -1 {
			missCounter++
			continue
		}
		pair := m.ArchiveCopy[j]

		dist := arrays.EuclideanDistance(pair.Fitness().ObjectiveValues, p.Fitness().ObjectiveValues)

		if dist <= m.BoundaryMinDistance {
			m.BoundaryPairs[j] = -1
			missCounter++
			continue
		}

		middlePoint := arrays.Middle(p.Genotype(), pair.Genotype())
		ind := MoeadIndividual{D: len(p.Genotype())}
		ind.SetGenotype(middlePoint)
		ind.Repair()
		ind.UpdateFitness(m.CMOP)

		if feasible(ind.Fitness()) {
			m.ArchiveCopy[j] = &ind
		} else {
			m.population[i] = &ind
		}
	}

	if missCounter <= m.historyCounter && missCounter > 0 {
		m.population = selectBinaryResult(m.ArchiveCopy, m.population, m.populationSize, m.BoundaryFeasibleSelectionProbability)
		m.binaryCompleted = true
		m.maxViolation = -1
		for _, p := range m.population {
			cv := constraintViolation(p.Fitness())
			if cv > m.maxViolation {
				m.maxViolation = cv
			}
		}
	}
	m.historyCounter = missCounter
}

func (m Moead) selectRandomPairs() []int {
	indices := make([]int, len(m.population))
	for i := range m.population {
		indices[i] = rand.Intn(len(m.archive))
	}
	return indices
}

func (m Moead) archiveCopy() []types.Individual {
	var arcCopy []types.Individual
	for _, i := range m.archive {
		arcCopy = append(arcCopy, i.Copy())
	}
	return arcCopy
}
