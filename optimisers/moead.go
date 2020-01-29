package optimisers

import (
	"math/rand"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
)

type Individual struct {
	D        int
	Genotype arrays.Vector
}

type Population struct {
	P           int
	D           int
	Individuals []Individual
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
}

func (p *Population) Initialise() {
	for i := 0; i < p.P; i++ {
		ind := Individual{D: p.D}
		ind.InitialiseRandom()
		p.Individuals = append(p.Individuals, ind)
	}
}

func (i *Individual) InitialiseRandom() {
	i.Genotype = arrays.Vector{Size: i.D}
	i.Genotype.Zeros()
	for j := 0; j < i.Genotype.Size; j++ {
		i.Genotype.Set(j, rand.Float64())
	}
}
