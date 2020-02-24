package testSuite3

import (
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func CIMOP1(x types.Genotype) types.Fitness {
	return types.Fitness{
		ObjectiveCount: 2,
		ObjectiveValues: []float64{
			objective1(x),
			objective2(x),
		},
		ConstraintCount: 1,
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrLessThanZero,
		},
		ConstraintValues: []float64{
			constraint1(x),
		},
	}
}

func CIMOP2(x types.Genotype) types.Fitness {
	return types.Fitness{
		ObjectiveCount: 2,
		ObjectiveValues: []float64{
			objective3(x),
			objective4(x),
		},
		ConstraintCount: 1,
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrLessThanZero,
		},
		ConstraintValues: []float64{
			constraint1(x),
		},
	}
}

func CIMOP3(x types.Genotype) types.Fitness {
	return types.Fitness{
		ObjectiveCount: 2,
		ObjectiveValues: []float64{
			objective5(x),
			objective6(x),
		},
		ConstraintCount: 1,
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrGreaterThanZero,
		},
		ConstraintValues: []float64{
			constraint1(x),
		},
	}
}

func CIMOP4(x types.Genotype) types.Fitness {
	return types.Fitness{
		ObjectiveCount: 2,
		ObjectiveValues: []float64{
			objective7(x),
			objective8(x),
		},
		ConstraintCount: 1,
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrGreaterThanZero,
		},
		ConstraintValues: []float64{
			constraint1(x),
		},
	}
}
