package cmop

import (
	"github.com/CRAB-LAB-NTNU/PPS-BS/testSuite"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

type lir struct {
	CMOP1  types.CMOP
	CMOP2  types.CMOP
	CMOP3  types.CMOP
	CMOP4  types.CMOP
	CMOP5  types.CMOP
	CMOP6  types.CMOP
	CMOP7  types.CMOP
	CMOP8  types.CMOP
	CMOP9  types.CMOP
	CMOP10 types.CMOP
	CMOP11 types.CMOP
	CMOP12 types.CMOP
	CMOP13 types.CMOP
	CMOP14 types.CMOP
}

func buildLirCmop1() types.CMOP {
	cmop1 := types.CMOP{
		ConstraintCount: 2, ObjectiveCount: 2,
		DecisionVariables: 30,
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrGreaterThanZero,
			types.EqualsOrGreaterThanZero,
		},
		Name: "LIR-CMOP1",
		TrueParetoFrontFilename: "arraydata/pf_data/CMOP1.dat",
	}
	cmop1.SetDecisionInterval(0, 1)
	cmop1.Evaluate = func(x types.Genotype) types.Fitness {

		f := initFitness(cmop1)

		f.ObjectiveValues = []float64{
			testSuite.Objective1(x),
			testSuite.Objective2(x),
		}

		f.ConstraintValues = []float64{
			testSuite.Constraint1(x, testSuite.Inner1),
			testSuite.Constraint1(x, testSuite.Inner2),
		}

		return f
	}
	return cmop1
}

func initFitness(cmop types.CMOP) types.Fitness {
	return types.Fitness{
		ObjectiveCount:  cmop.ObjectiveCount,
		ConstraintCount: cmop.ConstraintCount,
		ConstraintTypes: cmop.ConstraintTypes,
	}
}

func Lir() lir {
	return lir{
		CMOP1: buildLirCmop1(),
	}
}
