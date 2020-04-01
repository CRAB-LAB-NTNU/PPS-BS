package testsuites

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

var lirCmop1 = types.Cmop{
	ConstraintCount: 2, ObjectiveCount: 2,
	DecisionVariables: 30,
	ConstraintTypes: []types.ConstraintType{
		types.EqualsOrGreaterThanZero,
		types.EqualsOrGreaterThanZero,
	},
	Name:                    "LIR-CMOP1",
	TrueParetoFrontFilename: "paretoFrontData/lir/PF1.dat",
	DecisionInterval:        arrays.EqualInterval(30, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount: 2, ConstraintCount: 2,
			ConstraintTypes: []types.ConstraintType{
				types.EqualsOrGreaterThanZero,
				types.EqualsOrGreaterThanZero,
			},
		}
		fitness.ObjectiveValues = []float64{
			lirobjective1(x),
			lirobjective2(x),
		}
		fitness.ConstraintValues = []float64{
			lirconstraint1(x, lirinner1),
			lirconstraint1(x, lirinner2),
		}
		return fitness
	},
}

var lirCmop2 = types.Cmop{
	ConstraintCount: 2, ObjectiveCount: 2,
	DecisionVariables: 30,
	ConstraintTypes: []types.ConstraintType{
		types.EqualsOrGreaterThanZero,
		types.EqualsOrGreaterThanZero,
	},
	Name:                    "LIR-CMOP2",
	TrueParetoFrontFilename: "paretoFrontData/lir/PF2.dat",
	DecisionInterval:        arrays.EqualInterval(30, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		return types.Fitness{
			ObjectiveCount: 2, ConstraintCount: 2,
			ConstraintTypes: []types.ConstraintType{
				types.EqualsOrGreaterThanZero,
				types.EqualsOrGreaterThanZero,
			},
			ObjectiveValues: []float64{
				lirobjective1(x),
				lirobjective3(x),
			},
			ConstraintValues: []float64{
				lirconstraint1(x, lirinner1),
				lirconstraint1(x, lirinner2),
			},
		}
	},
}

var lirCmop3 = types.Cmop{
	ConstraintCount: 3, ObjectiveCount: 2,
	DecisionVariables: 30,
	ConstraintTypes: []types.ConstraintType{
		types.EqualsOrGreaterThanZero,
		types.EqualsOrGreaterThanZero,
		types.EqualsOrGreaterThanZero,
	},
	Name:                    "LIR-CMOP3",
	TrueParetoFrontFilename: "paretoFrontData/lir/PF3.dat",
	DecisionInterval:        arrays.EqualInterval(30, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		return types.Fitness{
			ObjectiveCount: 2, ConstraintCount: 3,
			ConstraintTypes: []types.ConstraintType{
				types.EqualsOrGreaterThanZero,
				types.EqualsOrGreaterThanZero,
				types.EqualsOrGreaterThanZero,
			},
			ObjectiveValues: []float64{
				lirobjective1(x),
				lirobjective2(x),
			},
			ConstraintValues: []float64{
				lirconstraint1(x, lirinner1),
				lirconstraint1(x, lirinner2),
				lirconstraint2(x),
			},
		}
	},
}

var lirCmop4 = types.Cmop{
	ConstraintCount: 3, ObjectiveCount: 2,
	DecisionVariables: 30,
	ConstraintTypes: []types.ConstraintType{
		types.EqualsOrGreaterThanZero,
		types.EqualsOrGreaterThanZero,
		types.EqualsOrGreaterThanZero,
	},
	Name:                    "LIR-CMOP4",
	TrueParetoFrontFilename: "paretoFrontData/lir/PF4.dat",
	DecisionInterval:        arrays.EqualInterval(30, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		return types.Fitness{
			ObjectiveCount:  2,
			ConstraintCount: 3,
			ConstraintTypes: []types.ConstraintType{
				types.EqualsOrGreaterThanZero,
				types.EqualsOrGreaterThanZero,
				types.EqualsOrGreaterThanZero,
			},
			ObjectiveValues: []float64{
				lirobjective1(x),
				lirobjective3(x),
			},
			ConstraintValues: []float64{
				lirconstraint1(x, lirinner1),
				lirconstraint1(x, lirinner2),
				lirconstraint2(x),
			},
		}
	},
}

var lirCmop5 = types.Cmop{
	ConstraintCount: 2, ObjectiveCount: 2,
	DecisionVariables: 30,
	ConstraintTypes: []types.ConstraintType{
		types.EqualsOrGreaterThanZero,
		types.EqualsOrGreaterThanZero,
	},
	Name:                    "LIR-CMOP5",
	TrueParetoFrontFilename: "paretoFrontData/lir/PF5.dat",
	DecisionInterval:        arrays.EqualInterval(30, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount:  2,
			ConstraintCount: 2,
			ConstraintTypes: []types.ConstraintType{
				types.EqualsOrGreaterThanZero,
				types.EqualsOrGreaterThanZero,
			},
		}
		p := []float64{1.6, 2.5}
		q := []float64{1.6, 2.5}
		a := []float64{2, 2}
		b := []float64{4, 8}
		f1 := lirobjective4(x)
		f2 := lirobjective5(x)
		c1 := lirconstraint3(x, p, q, a, b, 0, f1, f2)
		c2 := lirconstraint3(x, p, q, a, b, 1, f1, f2)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2}
		return fitness
	},
}

var lirCmop6 = types.Cmop{
	ConstraintCount: 2, ObjectiveCount: 2,
	DecisionVariables: 30,
	ConstraintTypes: []types.ConstraintType{
		types.EqualsOrGreaterThanZero,
		types.EqualsOrGreaterThanZero,
	},
	Name:                    "LIR-CMOP6",
	TrueParetoFrontFilename: "paretoFrontData/lir/PF6.dat",
	DecisionInterval:        arrays.EqualInterval(30, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount:  2,
			ConstraintCount: 2,
			ConstraintTypes: []types.ConstraintType{
				types.EqualsOrGreaterThanZero,
				types.EqualsOrGreaterThanZero,
			},
		}
		p := []float64{1.8, 2.8}
		q := []float64{1.8, 2.8}
		a := []float64{2, 2}
		b := []float64{8, 8}
		f1 := lirobjective4(x)
		f2 := lirobjective6(x)
		c1 := lirconstraint3(x, p, q, a, b, 0, f1, f2)
		c2 := lirconstraint3(x, p, q, a, b, 1, f1, f2)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2}
		return fitness
	},
}

var lirCmop7 = types.Cmop{
	ObjectiveCount: 2, ConstraintCount: 3,
	DecisionVariables:       30,
	ConstraintTypes:         []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
	Name:                    "LIR-CMOP7",
	TrueParetoFrontFilename: "paretoFrontData/lir/PF7.dat",
	DecisionInterval:        arrays.EqualInterval(30, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount: 2, ConstraintCount: 3,
			ConstraintTypes: []types.ConstraintType{
				types.EqualsOrGreaterThanZero,
				types.EqualsOrGreaterThanZero,
				types.EqualsOrGreaterThanZero,
			},
		}
		p := []float64{1.2, 2.25, 3.5}
		q := []float64{1.2, 2.25, 3.5}
		a := []float64{2, 2.5, 2.5}
		b := []float64{6, 12, 10}
		f1, f2 := lirobjective4(x), lirobjective5(x)
		c1 := lirconstraint3(x, p, q, a, b, 0, f1, f2)
		c2 := lirconstraint3(x, p, q, a, b, 1, f1, f2)
		c3 := lirconstraint3(x, p, q, a, b, 2, f1, f2)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2, c3}
		return fitness
	},
}

var lirCmop8 = types.Cmop{
	ObjectiveCount: 2, ConstraintCount: 3,
	DecisionVariables:       30,
	ConstraintTypes:         []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
	Name:                    "LIR-CMOP8",
	TrueParetoFrontFilename: "paretoFrontData/lir/PF8.dat",
	DecisionInterval:        arrays.EqualInterval(30, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount: 2, ConstraintCount: 3,
			ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		}
		p := []float64{1.2, 2.25, 3.5}
		q := []float64{1.2, 2.25, 3.5}
		a := []float64{2, 2.5, 2.5}
		b := []float64{6, 12, 10}
		f1, f2 := lirobjective4(x), lirobjective6(x)
		c1 := lirconstraint3(x, p, q, a, b, 0, f1, f2)
		c2 := lirconstraint3(x, p, q, a, b, 1, f1, f2)
		c3 := lirconstraint3(x, p, q, a, b, 2, f1, f2)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2, c3}
		return fitness
	},
}

var lirCmop9 = types.Cmop{
	ObjectiveCount: 2, ConstraintCount: 2,
	DecisionVariables:       30,
	ConstraintTypes:         []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
	Name:                    "LIR-CMOP9",
	TrueParetoFrontFilename: "paretoFrontData/lir/PF9.dat",
	DecisionInterval:        arrays.EqualInterval(30, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount: 2, ConstraintCount: 2,
			ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		}
		p := []float64{1.4}
		q := []float64{1.4}
		a := []float64{1.5}
		b := []float64{6.0}
		f1, f2 := lirobjective7(x), lirobjective8(x)
		c1 := lirconstraint3(x, p, q, a, b, 0, f1, f2)
		c2 := lirconstraint4(x, f1, f2, 2)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2}
		return fitness
	},
}

var lirCmop10 = types.Cmop{
	ObjectiveCount: 2, ConstraintCount: 2,
	DecisionVariables:       30,
	ConstraintTypes:         []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
	Name:                    "LIR-CMOP10",
	TrueParetoFrontFilename: "paretoFrontData/lir/PF10.dat",
	DecisionInterval:        arrays.EqualInterval(30, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount: 2, ConstraintCount: 2,
			ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		}
		p := []float64{1.1}
		q := []float64{1.2}
		a := []float64{2.0}
		b := []float64{4.0}
		f1, f2 := lirobjective7(x), lirobjective9(x)
		c1 := lirconstraint3(x, p, q, a, b, 0, f1, f2)
		c2 := lirconstraint4(x, f1, f2, 1)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2}
		return fitness
	},
}

var lirCmop11 = types.Cmop{
	ObjectiveCount: 2, ConstraintCount: 2,
	DecisionVariables:       30,
	ConstraintTypes:         []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
	Name:                    "LIR-CMOP11",
	TrueParetoFrontFilename: "paretoFrontData/lir/PF11.dat",
	DecisionInterval:        arrays.EqualInterval(30, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount: 2, ConstraintCount: 2,
			ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		}
		p := []float64{1.2}
		q := []float64{1.2}
		a := []float64{1.5}
		b := []float64{5.0}
		f1, f2 := lirobjective7(x), lirobjective9(x)
		c1 := lirconstraint3(x, p, q, a, b, 0, f1, f2)
		c2 := lirconstraint4(x, f1, f2, 2.1)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2}
		return fitness
	},
}

var lirCmop12 = types.Cmop{
	ObjectiveCount: 2, ConstraintCount: 2,
	DecisionVariables:       30,
	ConstraintTypes:         []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
	Name:                    "LIR-CMOP12",
	TrueParetoFrontFilename: "paretoFrontData/lir/PF12.dat",
	DecisionInterval:        arrays.EqualInterval(30, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount: 2, ConstraintCount: 2,
			ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		}
		p := []float64{1.6}
		q := []float64{1.6}
		a := []float64{1.5}
		b := []float64{6.0}
		f1, f2 := lirobjective7(x), lirobjective8(x)
		c1 := lirconstraint3(x, p, q, a, b, 0, f1, f2)
		c2 := lirconstraint4(x, f1, f2, 2.5)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2}
		return fitness
	},
}

var lirCmop13 = types.Cmop{
	ObjectiveCount: 3, ConstraintCount: 2,
	ConstraintTypes:         []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
	DecisionVariables:       30,
	Name:                    "LIR-CMOP13",
	TrueParetoFrontFilename: "paretoFrontData/lir/PF13.dat",
	DecisionInterval:        arrays.EqualInterval(30, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount: 3, ConstraintCount: 2,
			ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		}
		f1, f2, f3 := lirobjective10(x), lirobjective11(x), lirobjective12(x)
		g := math.Pow(f1, 2) + math.Pow(f2, 2) + math.Pow(f3, 2)
		c1, c2 := lirconstraint5(x, g, 9, 4), lirconstraint5(x, g, 3.61, 3.24)
		fitness.ObjectiveValues = []float64{f1, f2, f3}
		fitness.ConstraintValues = []float64{c1, c2}
		return fitness
	},
}

var lirCmop14 = types.Cmop{
	ObjectiveCount: 3, ConstraintCount: 3,
	ConstraintTypes:         []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
	DecisionVariables:       30,
	Name:                    "LIR-CMOP14",
	TrueParetoFrontFilename: "paretoFrontData/lir/PF14.dat",
	DecisionInterval:        arrays.EqualInterval(30, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount: 3, ConstraintCount: 3,
			ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		}
		f1, f2, f3 := lirobjective10(x), lirobjective11(x), lirobjective12(x)
		g := math.Pow(f1, 2) + math.Pow(f2, 2) + math.Pow(f3, 2)
		c1, c2, c3 := lirconstraint5(x, g, 9, 4), lirconstraint5(x, g, 3.61, 3.24), lirconstraint5(x, g, 3.0625, 2.56)
		fitness.ObjectiveValues = []float64{f1, f2, f3}
		fitness.ConstraintValues = []float64{c1, c2, c3}
		return fitness
	},
}
