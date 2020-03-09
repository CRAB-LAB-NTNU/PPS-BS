package optimisers

import (
	"math"
	"math/rand"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
	"github.com/CRAB-LAB-NTNU/PPS-BS/utils"
)

/*MoeadIndividual is a struct containing information about an individual in the population
of an evolutionary algorithm.
*/
type MoeadIndividual struct {
	Cmop     types.CMOP
	genotype types.Genotype
	fitness  types.Fitness
}

func (i *MoeadIndividual) SetGenotype(g []float64) {
	i.genotype = g
}

func (i *MoeadIndividual) SetFitness(f types.Fitness) {
	i.fitness = f
}

func (i MoeadIndividual) Copy() types.Individual {
	ind := MoeadIndividual{
		Cmop:     i.Cmop,
		fitness:  i.Fitness(),
		genotype: i.Genotype(),
	}
	return &ind
}

/*Genotype returns the genotype of the individual
 */
func (i MoeadIndividual) Genotype() types.Genotype {
	return i.genotype
}

/*Fitness returnes the fitness value of the individual if it has been calculated
 */
func (i MoeadIndividual) Fitness() types.Fitness {
	return i.fitness
}

/*UpdateFitness updates the fitness value of an individual
 */
func (i *MoeadIndividual) UpdateFitness() types.Fitness {
	i.fitness = i.Cmop.Evaluate(i.Genotype())
	return i.Fitness()
}

func (i MoeadIndividual) D() int {
	return i.Cmop.DecisionVariables
}

/*Initialise initialises the individuals genotype with random floats in the range [0,1]
 */
func (i *MoeadIndividual) Initialise() {
	i.genotype = make([]float64, i.D())
	for j := range i.genotype {
		interval := i.Cmop.DecisionInterval[j]
		i.genotype[j] = utils.RandomFloat64Range(interval[0], interval[1])
	}
	i.UpdateFitness()
}

func (ind *MoeadIndividual) Repair() {
	for i, gene := range ind.genotype {
		min, max := ind.Cmop.DecisionInterval[i][0], ind.Cmop.DecisionInterval[i][1]
		if gene > max {
			ind.genotype[i] = max
		} else if gene < min {
			ind.genotype[i] = min
		}
	}
}

/*PolynomialMutation performs the mutation described in the
paper https://reader.elsevier.com/reader/sd/pii/S0045782599003898
PPS describes a mutation probability of 1/n where n => length of genotype.
We don't iterate but make a 1/n dice roll to check if we're mutating a single alele.
*/
func (i *MoeadIndividual) PolynomialMutation(m float64) {
	if m < 0 {
		panic("m needs to be non-negative")
	}

	for pos, allele := range i.genotype {
		if rand.Float64() > 1/float64(i.D()) {
			continue
		}
		u := rand.Float64()
		delta := math.Min((allele-0), (1-allele)) / (1 - 0)
		var gresk float64
		if u <= 0.5 {
			gresk = math.Pow(2*u+math.Pow((1-2*u)*(1-delta), (m+1)), (1/(m+1))) - 1
		} else {
			gresk = 1 - math.Pow(2*(1-u)+math.Pow(2*(u-0.5)*(1-delta), (m+1)), (1/(m+1)))
		}
		i.genotype[pos] += gresk
	}
}
