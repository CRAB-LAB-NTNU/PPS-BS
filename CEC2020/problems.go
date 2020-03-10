package CEC2020

import (
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func CEC_CMOP1(x types.Genotype) types.Fitness {
	fitness := types.Fitness{
		ObjectiveCount:  2,
		ConstraintCount: 3,
	}

	fitness.ConstraintTypes = []types.ConstraintType{
		types.EqualsOrLessThanZero,
		types.EqualsOrLessThanZero,
		types.EqualsOrLessThanZero,
	}
	return fitness
}
