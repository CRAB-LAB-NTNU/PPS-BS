package optimisers

import (
	"math/rand"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func (m *Moead) PushProblems(j int, y types.Individual) bool {
	xF := m.population[j].Fitness()
	yF := y.Fitness()

	xS := tchebycheff(xF.ObjectiveValues, m.IdealPoint, m.Weights[j])
	yS := tchebycheff(yF.ObjectiveValues, m.IdealPoint, m.Weights[j])

	if yS <= xS {
		m.population[j] = y
		return true
	}
	return false
}

func (m *Moead) PullProblems(j int, y types.Individual, eps float64) bool {
	xF := m.population[j].Fitness()
	yF := y.Fitness()
	xCV := maximumConstraintViolation(xF)
	yCV := maximumConstraintViolation(yF)
	xS := tchebycheff(xF.ObjectiveValues, m.IdealPoint, m.Weights[j])
	yS := tchebycheff(yF.ObjectiveValues, m.IdealPoint, m.Weights[j])

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
	} else {
		hood := make([]int, m.populationSize)
		for i := range hood {
			hood[i] = i
		}
		return hood
	}
}
