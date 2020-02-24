package testsuite2

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func IMB11(x types.Genotype) types.Fitness {
	return types.Fitness{
		ObjectiveCount: 2,
		ObjectiveValues: []float64{
			objective1(x),
			objective2(x),
		},
		ConstraintCount: 1,
		ConstraintValues: []float64{
			math.SmallestNonzeroFloat64 - constraint1(x),
		},
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrGreaterThanZero,
		},
	}
}

func IMB12(x types.Genotype) types.Fitness {
	return types.Fitness{
		ObjectiveCount: 2,
		ObjectiveValues: []float64{
			objective1(x),
			objective3(x),
		},
		ConstraintCount: 1,
		ConstraintValues: []float64{
			math.SmallestNonzeroFloat64 - constraint2(x),
		},
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrGreaterThanZero,
		},
	}
}

func IMB13(x types.Genotype) types.Fitness {
	return types.Fitness{
		ObjectiveCount: 2,
		ObjectiveValues: []float64{
			objective1(x),
			objective4(x),
		},
		ConstraintCount: 1,
		ConstraintValues: []float64{
			math.SmallestNonzeroFloat64 - constraint3(x),
		},
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrGreaterThanZero,
		},
	}
}
