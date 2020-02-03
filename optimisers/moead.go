package optimisers

import (
	"math/rand"

	"github.com/CRAB-LAB-NTNU/PPS-BS/testSuite"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
)

type Individual struct {
	D        int
	Genotype arrays.Vector
	fitness  testSuite.Fitness
	CMOP     func(arrays.Vector) testSuite.Fitness
}

type Population struct {
	P           int
	D           int
	Individuals []Individual
	IdealPoint  []float64
}

type Moead struct {
	EP                 []Individual
	Population         Population
	O, T, N, D         int
	Weights            []arrays.Vector
	WeightNeigbourhood [][]int
}

func (m *Moead) Initialise() {
	m.Weights = arrays.UniformDistributedVectors(m.O, m.N)

	for i := range m.Weights {
		m.WeightNeigbourhood = append(m.WeightNeigbourhood, arrays.NearestNeighbour(m.Weights, i, m.T))
	}
	m.Population = Population{P: m.N, D: m.D}
	m.Population.Initialise()
	m.Population.CalculateIdealPoints()
}

func (p *Population) Initialise() {
	for i := 0; i < p.P; i++ {
		ind := Individual{D: p.D}
		ind.InitialiseRandom()
		p.Individuals = append(p.Individuals, ind)
	}
}

func (p *Population) CalculateIdealPoints() {
	// TODO
	p.IdealPoint = []float64{0, 8, 0, 1}
}

func (i *Individual) InitialiseRandom() {
	i.Genotype = arrays.Vector{Size: i.D}
	i.Genotype.Zeros()
	for j := 0; j < i.Genotype.Size; j++ {
		i.Genotype.Set(j, rand.Float64())
	}
	i.UpdateFitness()
}

func (i *Individual) UpdateFitness() {
	i.fitness = i.CMOP(i.Genotype)
}

func (i Individual) Fitness() testSuite.Fitness {
	return i.fitness
}

func (m *Moead) Crossover([]Individual) []Individual {
	//Ronk.
	return []Individual{}
}

func (m *Moead) Evolve() {
	newPopulation := Population{P: m.Population.P, D: m.Population.D}
	for i := 0; i < m.N; i++ {
		x, y := rand.Intn(len(m.WeightNeigbourhood[i])), rand.Intn(len(m.WeightNeigbourhood[i]))
		offSpring := m.Crossover([]Individual{m.Population.Individuals[x], m.Population.Individuals[y]})
		newPopulation.Individuals = append(newPopulation.Individuals, offSpring[0])
		// Original paper has a repair function or improvement heuristic at this point in the algorithm.
		// The original paper uses a repair function for the Knapsack problem.
		// Which i @THINK@ is their way of doing some constraint handling.
		// As an improvement heuristic is problem specific, i guess the mutation step can be used for this purpose.
		// We might need a Mutation Function aswell as a Repair function.
		// The PPS paper uses a more traditional DE variant by comparing in the first phase.

	}
}
