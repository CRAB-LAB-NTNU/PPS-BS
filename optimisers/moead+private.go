package optimisers

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/chm"
	"github.com/CRAB-LAB-NTNU/PPS-BS/stages"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

/*crossover uses the traditional DE operator to generate a new individual
Our MOEA/D is implemented with a single offspring and parents picked from the
weight neighbourhood.
*/
func (m *Moead) crossover(p, x, y types.Individual) types.Individual {

	child := MoeadIndividual{
		D:        m.DecisionVariables,
		genotype: make([]float64, m.DecisionVariables),
	}

	for i := range child.Genotype() {
		if rand.Float64() < m.Cr {
			child.Genotype()[i] = p.Genotype()[i] + m.F*(x.Genotype()[i]-y.Genotype()[i])
		} else {
			child.Genotype()[i] = p.Genotype()[i]
		}
	}

	child.PolynomialMutation(m.DistributionIndex)
	child.Repair()
	child.UpdateFitness(m.cmop)
	return &child
}

func (m *Moead) updateIdealPoint(offspring types.Individual) {
	f := offspring.Fitness()
	for pos, val := range f.ObjectiveValues {
		m.idealPoint[pos] = math.Min(m.idealPoint[pos], val)
	}
}

func (m *Moead) updateMaxConstraintViolation(offspring types.Individual) {
	f := offspring.Fitness()
	m.maxViolation = math.Max(m.maxViolation, f.TotalViolation())
}

func (m *Moead) updatePopulation(hood []int, offspring types.Individual, replace func(int, types.Individual) bool) {
	c := 0
	for c < m.Nr && len(hood) > 0 {
		j := rand.Intn(len(hood))
		replaced := replace(hood[j], offspring)
		if replaced {
			c++
		}
		hood = arrays.Remove(hood, j)
	}
}

// Used during the push phase of the algorithm when constraints are ignored
func (m *Moead) replaceIgnoringConstraints(p int, o types.Individual) bool {
	pF := m.population[p].Fitness()
	oF := o.Fitness()

	pS := tchebycheff(pF.ObjectiveValues, m.idealPoint, m.Weights[p])
	oS := tchebycheff(oF.ObjectiveValues, m.idealPoint, m.Weights[p])

	if oS <= pS {
		m.population[p] = o
		return true
	}
	return false
}

// Used when constraints are not ignored
func (m *Moead) replaceWithConstraints(p int, o types.Individual) bool {
	pF := m.population[p].Fitness()
	oF := o.Fitness()
	pCV := m.chm.Violation(m.generation, pF)
	oCV := m.chm.Violation(m.generation, oF)
	pS := tchebycheff(pF.ObjectiveValues, m.idealPoint, m.Weights[p])
	oS := tchebycheff(oF.ObjectiveValues, m.idealPoint, m.Weights[p])

	if oCV <= 0 && pCV <= 0 {
		if oS <= pS {
			m.population[p] = o
			return true
		}
	} else if oCV == pCV {
		if oS <= pS {
			m.population[p] = o
			return true
		}
	} else if oCV < pCV {
		m.population[p] = o
		return true
	}

	return false
}

func (m Moead) selectHood(pr float64, i int) []int {
	if rand.Float64() < pr {
		hood := make([]int, m.T)
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

func (m Moead) selectRandomPairs() []int {
	indices := make([]int, len(m.population))
	for i := range m.population {
		indices[i] = rand.Intn(len(m.archive))
	}
	return indices
}

func (m Moead) copyArchive() []types.Individual {
	var arcCopy []types.Individual
	for _, i := range m.archive {
		arcCopy = append(arcCopy, i.Copy())
	}
	return arcCopy
}

// Binary Specific Methods

func (m *Moead) evolveBinary(stage types.Stage) {
	binary, ok := stage.(*stages.Binary)
	if !ok {
		panic("Could not assert binary stage")
	}
	if m.skipBinary() {
		fmt.Println("No feasible individuals in the archive - Skipping binary search")
		stage.SetOver()
	} else {
		if !m.hasBinaryPairs() {
			m.binaryPairs = m.selectRandomPairs()
			m.archiveCopy = m.copyArchive()
		}
		m.boundarySearch(binary)
		m.generation++
	}
}

func (m Moead) skipBinary() bool {
	return len(m.archive) == 0
}

func (m Moead) hasBinaryPairs() bool {
	return len(m.binaryPairs) > 0
}

func (m *Moead) boundarySearch(binary *stages.Binary) {
	missCounter := 0
	for i, p := range m.population {
		m.fnEval++
		j := m.binaryPairs[i]
		if j == -1 {
			missCounter++
			continue
		}
		pair := m.archiveCopy[j]

		dist := arrays.EuclideanDistance(pair.Fitness().ObjectiveValues, p.Fitness().ObjectiveValues)

		if dist <= binary.MinDistance() {
			m.binaryPairs[j] = -1
			missCounter++
			continue
		}

		middlePoint := arrays.Middle(p.Genotype(), pair.Genotype())
		ind := MoeadIndividual{D: len(p.Genotype())}
		ind.SetGenotype(middlePoint)
		ind.Repair()
		ind.UpdateFitness(m.cmop)

		if feasible(ind.Fitness()) {
			m.archiveCopy[j] = &ind
		} else {
			m.population[i] = &ind
		}
	}
	if missCounter <= m.historyCounter && missCounter > 0 {
		m.population = selectBinaryResult(m.archiveCopy, m.population, m.populationSize, binary.Fcp())
		fmt.Println("Setting binary stage over")
		binary.SetOver()
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

// CHM
func (m *Moead) updateCHM() {

	// We try to cast to r2s to see if that is the constraint handling method used.
	// This is because we have to check for active constraints
	r2s, ok := m.chm.(*chm.R2S)
	if ok {
		m.determineActiveConstraints(r2s)
		r2s.Update(m.generation, float64(m.fnEval))
		m.chm = r2s
		return
	}
	m.chm.Update(m.generation, m.FeasibleRatio())

}

// R2S specific methods

func (m *Moead) determineActiveConstraints(r2s *chm.R2S) {
	rankedPopulation := biooperators.FastNonDominatedSort(m.population)
	//TODO determine a good selection mechanism for the individual
	randomIndex := rand.Intn(len(rankedPopulation[0]))
	randomBest := rankedPopulation[0][randomIndex]

	r2s.ACD(m.generation, m.fnEval, randomBest.Fitness())
}
