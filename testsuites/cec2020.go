package testsuites

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

var cecCmop1 = types.Cmop{
	ConstraintCount:         3,
	ObjectiveCount:          2,
	DecisionVariables:       25,
	ConstraintTypes:         []types.ConstraintType{types.EqualsOrLessThanZero, types.EqualsOrLessThanZero, types.EqualsOrLessThanZero},
	Name:                    "CEC2020-CMOP1",
	TrueParetoFrontFilename: "paretoFrontData/cec2020/PF1.dat",
	DecisionInterval:        arrays.EqualInterval(25, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount: 2, ConstraintCount: 3,
			ConstraintTypes: []types.ConstraintType{types.EqualsOrLessThanZero, types.EqualsOrLessThanZero, types.EqualsOrLessThanZero},
		}
		g := cecinner1(x)
		f1 := x[0] * g
		f2 := g * math.Sqrt(1-math.Pow(f1/g, 2))
		l := math.Atan(f2 / f1)
		c1 := math.Pow(f1, 2) + math.Pow(f2, 2) - math.Pow(1.7-0.2*math.Sin(2*l), 2)
		c2 := math.Pow(1+0.5*math.Sin(6*math.Pow(0.5*math.Pi-2*math.Abs(l-0.25*math.Pi), 3)), 2) - math.Pow(f1, 2) - math.Pow(f2, 2)
		c3 := math.Pow(1-0.45*math.Sin(6*math.Pow(0.5*math.Pi-2*math.Abs(l-0.25*math.Pi), 3)), 2) - math.Pow(f1, 2) - math.Pow(f2, 2)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2, c3}
		return fitness
	},
}

var cecCmop2 = types.Cmop{
	ConstraintCount:         1,
	ObjectiveCount:          2,
	DecisionVariables:       25,
	ConstraintTypes:         []types.ConstraintType{types.EqualsOrLessThanZero},
	Name:                    "CEC2020-CMOP2",
	TrueParetoFrontFilename: "paretoFrontData/cec2020/PF2.dat",
	DecisionInterval:        arrays.EqualInterval(25, 0, 1.1),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount: 2, ConstraintCount: 1,
			ConstraintTypes: []types.ConstraintType{types.EqualsOrLessThanZero},
		}
		g := cecinner2(x)
		f1 := x[0] * g
		f2 := g * math.Sqrt(math.Pow(1.1, 2)-math.Pow(f1/g, 2))
		tan := math.Pow(math.Atan(f2/f1), 4)
		l := math.Pow(math.Cos(6*tan), 10)
		c1 := math.Pow(f1/(1+0.15*l), 2) + math.Pow(f2/(1+0.75*l), 2) - 1
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1}
		return fitness
	},
}

var cecCmop3 = types.Cmop{
	ConstraintCount:         3,
	ObjectiveCount:          2,
	DecisionVariables:       25,
	ConstraintTypes:         []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrLessThanZero, types.EqualsOrLessThanZero},
	Name:                    "CEC2020-CMOP3",
	TrueParetoFrontFilename: "paretoFrontData/cec2020/PF3.dat",
	DecisionInterval:        arrays.EqualInterval(25, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount: 2, ConstraintCount: 3,
			ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrLessThanZero, types.EqualsOrLessThanZero},
		}
		g := cecinner2(x)
		f1 := g * math.Pow(x[0], float64(len(x)))
		f2 := g * (1 - math.Pow(f1/g, 2))
		c1 := (2 - 4*math.Pow(f1, 2) - f2) * (2 - 8*math.Pow(f1, 2) - f2)
		c2 := (2 - 2*math.Pow(f1, 2) - f2) * (2 - 16*math.Pow(f1, 2) - f2)
		c3 := (1 - math.Pow(f1, 2) - f2) * (1.2 - 1.2*math.Pow(f1, 2) - f2)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2, c3}
		return fitness
	},
}

var cecCmop4 = types.Cmop{
	ConstraintCount:         2,
	ObjectiveCount:          2,
	DecisionVariables:       25,
	ConstraintTypes:         []types.ConstraintType{types.EqualsOrLessThanZero, types.EqualsOrGreaterThanZero},
	Name:                    "CEC2020-CMOP4",
	TrueParetoFrontFilename: "paretoFrontData/cec2020/PF4.dat",
	DecisionInterval:        arrays.EqualInterval(25, 0, 1.5),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount: 2, ConstraintCount: 2,
			ConstraintTypes: []types.ConstraintType{types.EqualsOrLessThanZero, types.EqualsOrGreaterThanZero},
		}
		g := cecinner2(x)
		f1 := g * x[0]
		f2 := g * (5 - math.Exp(f1/g) - 0.5*math.Abs(math.Sin((3*math.Pi*f1)/g)))
		c1Prod1 := (5 - math.Exp(f1) - 0.5*math.Sin(3*math.Pi*f1) - f2)
		c1Prod2 := (5 - (1 + 0.4*f1) - 0.5*math.Sin(3*math.Pi*f1) - f2)
		c1 := c1Prod1 * c1Prod2
		c2Prod1 := (5 - (1 + f1 + 0.5*math.Pow(f1, 2)) - 0.5*math.Sin(3*math.Pi*f1) - f2)
		c2Prod2 := (5 - (1 + 0.7*f1) - 0.5*math.Sin(3*math.Pi*f1) - f2)
		c2 := c2Prod1 * c2Prod2
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2}
		return fitness
	},
}

var cecCmop5 = types.Cmop{
	ConstraintCount:         1,
	ObjectiveCount:          2,
	DecisionVariables:       25,
	ConstraintTypes:         []types.ConstraintType{types.EqualsOrLessThanZero},
	Name:                    "CEC2020-CMOP5",
	TrueParetoFrontFilename: "paretoFrontData/cec2020/PF5.dat",
	DecisionInterval:        arrays.EqualInterval(25, 0, 1),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ConstraintCount: 1, ObjectiveCount: 2,
			ConstraintTypes: []types.ConstraintType{types.EqualsOrLessThanZero},
		}
		g := cecinner1(x)
		f1 := g * x[0]
		f2 := g * (1 - math.Pow(f1/g, 0.6))
		T1 := (1 - 0.64*math.Pow(f1, 2) - f2) * (1 - 0.36*math.Pow(f1, 2) - f2)
		T2 := (math.Pow(1.35, 2) - math.Pow(f1+0.35, 2) - f2) * (math.Pow(1.15, 2) - math.Pow(f1+0.15, 2) - f2)
		c1 := math.Min(T1, T2)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1}
		return fitness
	},
}

var cecCmop6 = types.Cmop{
	ConstraintCount:   4,
	ObjectiveCount:    2,
	DecisionVariables: 25,
	ConstraintTypes: []types.ConstraintType{
		types.EqualsOrGreaterThanZero,
		types.EqualsOrLessThanZero,
		types.EqualsOrGreaterThanZero,
		types.EqualsOrLessThanZero,
	},
	Name:                    "CEC2020-CMOP6",
	TrueParetoFrontFilename: "paretoFrontData/cec2020/PF6.dat",
	DecisionInterval:        arrays.EqualInterval(25, 0, math.Sqrt(2)),
	Evaluate: func(x types.Genotype) types.Fitness {
		fitness := types.Fitness{
			ObjectiveCount: 2, ConstraintCount: 4,
			ConstraintTypes: []types.ConstraintType{
				types.EqualsOrGreaterThanZero,
				types.EqualsOrLessThanZero,
				types.EqualsOrGreaterThanZero,
				types.EqualsOrLessThanZero,
			},
		}
		g := cecinner3(x)
		f1 := g * x[0]
		// Because of rouding errors, (f1/g)^2 will ocasionally
		// return 2.000000000000001.
		// This makes 2 - (f1/g)^2 < 0 and the square root returns NaN.
		// We KNOW (f1/g)^2 <= 2 because of the substitution of f1:
		// (f1/g)^2 = ((g * x[0]) / g)^2 = x[0]^2
		// and all X is in the range [0,sqrt(2)].
		// The below truncation is therefore valid.
		// You can also check this by
		// var test bool = 2.0 == math.Pow(math.Sqrt(2.0), 2)
		f2 := g * math.Sqrt(2-math.Min(math.Pow(f1/g, 2), 2))

		c1 := (3 - math.Pow(f1, 2) - f2) * (3 - 2*math.Pow(f1, 2) - f2)
		c2 := (3 - 0.625*math.Pow(f1, 2) - f2) * (3 - 7*math.Pow(f1, 2) - f2)
		c3 := (1.62 - 0.18*math.Pow(f1, 2) - f2) * (1.125 - 0.125*math.Pow(f1, 2) - f2)
		c4 := (2.07 - 0.23*math.Pow(f1, 2) - f2) * (0.63 - 0.07*math.Pow(f1, 2) - f2)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2, c3, c4}
		return fitness
	},
}
