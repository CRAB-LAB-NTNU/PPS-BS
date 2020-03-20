package cmops

import "github.com/CRAB-LAB-NTNU/PPS-BS/types"

func initFitness(cmop types.Cmop) types.Fitness {
	return types.Fitness{
		ObjectiveCount:  cmop.ObjectiveCount,
		ConstraintCount: cmop.ConstraintCount,
		ConstraintTypes: cmop.ConstraintTypes,
	}
}
