package cmop

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

type cec2020 struct {
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
}

func buildCecCmop1() types.CMOP {
	cmop := types.CMOP{
		ConstraintCount:   3,
		ObjectiveCount:    2,
		DecisionVariables: 25,
		ConstraintTypes:   []types.ConstraintType{types.EqualsOrLessThanZero, types.EqualsOrLessThanZero, types.EqualsOrLessThanZero},
		Name:              "CEC2020-CMOP1",
		TrueParetoFrontFilename: "arraydata/cec2020/PF1.dat",
	}

	cmop.SetDecisionInterval(0, 1)
	cmop.Evaluate = func(x types.Genotype) types.Fitness {
		fitness := initFitness(cmop)

		g := g1(x)
		f1 := x[0] * g
		f2 := g * math.Sqrt(1-math.Pow(f1/g, 2))

		l := math.Atan(f2 / f1)
		c1 := math.Pow(f1, 2) + math.Pow(f2, 2) - math.Pow(1.7-0.2*math.Sin(2*l), 2)
		c2 := math.Pow(1+0.5*math.Sin(6*math.Pow(0.5*math.Pi-2*math.Abs(l-0.25*math.Pi), 3)), 2) - math.Pow(f1, 2) - math.Pow(f2, 2)
		c3 := math.Pow(1-0.45*math.Sin(6*math.Pow(0.5*math.Pi-2*math.Abs(l-0.25*math.Pi), 3)), 2) - math.Pow(f1, 2) - math.Pow(f2, 2)

		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2, c3}
		return fitness
	}

	return cmop
}

func buildCecCmop2() types.CMOP {
	cmop := types.CMOP{
		ConstraintCount:   1,
		ObjectiveCount:    2,
		DecisionVariables: 25,
		ConstraintTypes:   []types.ConstraintType{types.EqualsOrLessThanZero},
		Name:              "CEC2020-CMOP2",
		TrueParetoFrontFilename: "arraydata/cec2020/PF2.dat",
	}
	cmop.SetDecisionInterval(0, 1.1)
	cmop.Evaluate = func(x types.Genotype) types.Fitness {
		fitness := initFitness(cmop)
		g := G2(x)

		f1 := x[0] * g
		f2 := g * math.Sqrt(math.Pow(1.1, 2)-math.Pow(f1/g, 2))

		tan := math.Pow(math.Atan(f2/f1), 4)
		l := math.Pow(math.Cos(6*tan), 10)

		c1 := cmop2Constraint(f1, f2, l)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1}
		return fitness
	}
	return cmop
}

func buildCecCmop3() types.CMOP {
	cmop := types.CMOP{
		ConstraintCount:   3,
		ObjectiveCount:    2,
		DecisionVariables: 25,
		ConstraintTypes:   []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrLessThanZero, types.EqualsOrLessThanZero},
		Name:              "CEC2020-CMOP3",
		TrueParetoFrontFilename: "arraydata/cec2020/PF3.dat",
	}
	cmop.SetDecisionInterval(0, 1)
	cmop.Evaluate = func(x types.Genotype) types.Fitness {
		fitness := initFitness(cmop)
		g := G2(x)

		f1 := g * math.Pow(x[0], float64(len(x)))
		f2 := g * (1 - math.Pow(f1/g, 2))

		c1 := (2 - 4*math.Pow(f1, 2) - f2) * (2 - 8*math.Pow(f1, 2) - f2)
		c2 := (2 - 2*math.Pow(f1, 2) - f2) * (2 - 16*math.Pow(f1, 2) - f2)
		c3 := (1 - math.Pow(f1, 2) - f2) * (1.2 - 1.2*math.Pow(f1, 2) - f2)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2, c3}
		return fitness
	}
	return cmop
}

func buildCecCmop4() types.CMOP {
	cmop := types.CMOP{
		ConstraintCount:   2,
		ObjectiveCount:    2,
		DecisionVariables: 25,
		ConstraintTypes:   []types.ConstraintType{types.EqualsOrLessThanZero, types.EqualsOrGreaterThanZero},
		Name:              "CEC2020-CMOP4",
		TrueParetoFrontFilename: "arraydata/cec2020/PF4.dat",
	}
	cmop.SetDecisionInterval(0, 1.5)
	cmop.Evaluate = func(x types.Genotype) types.Fitness {
		fitness := initFitness(cmop)
		g := G2(x)

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
	}
	return cmop
}

func buildCecCmop5() types.CMOP {
	cmop := types.CMOP{
		ConstraintCount:   1,
		ObjectiveCount:    2,
		DecisionVariables: 25,
		ConstraintTypes:   []types.ConstraintType{types.EqualsOrLessThanZero},
		Name:              "CEC2020-CMOP5",
		TrueParetoFrontFilename: "arraydata/cec2020/PF5.dat",
	}
	cmop.SetDecisionInterval(0, 1)
	cmop.Evaluate = func(x types.Genotype) types.Fitness {
		fitness := initFitness(cmop)
		g := g1(x)

		f1 := g * x[0]
		f2 := g * (1 - math.Pow(f1/g, 0.6))
		T1 := (1 - 0.64*math.Pow(f1, 2) - f2) * (1 - 0.36*math.Pow(f1, 2) - f2)
		T2 := (math.Pow(1.35, 2) - math.Pow(f1+0.35, 2) - f2) * (math.Pow(1.15, 2) - math.Pow(f1+0.15, 2) - f2)
		c1 := math.Min(T1, T2)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1}
		return fitness
	}
	return cmop
}

func buildCecCmop6() types.CMOP {
	cmop := types.CMOP{
		ConstraintCount:   4,
		ObjectiveCount:    2,
		DecisionVariables: 25,
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrGreaterThanZero,
			types.EqualsOrLessThanZero,
			types.EqualsOrGreaterThanZero,
			types.EqualsOrLessThanZero,
		},
		Name: "CEC2020-CMOP6",
		TrueParetoFrontFilename: "arraydata/cec2020/PF6.dat",
	}
	cmop.SetDecisionInterval(0, math.Sqrt(2))
	cmop.Evaluate = func(x types.Genotype) types.Fitness {
		fitness := initFitness(cmop)
		g := g3(x)
		f1 := g * x[0]
		f2 := g * math.Sqrt(2-math.Pow(f1/g, 2))
		c1 := (3 - math.Pow(f1, 2) - f2) * (3 - 2*math.Pow(f1, 2) - f2)
		c2 := (3 - 0.625*math.Pow(f1, 2) - f2) * (3 - 7*math.Pow(f1, 2) - f2)
		c3 := (1.62 - 0.18*math.Pow(f1, 2) - f2) * (1.125 - 0.125*math.Pow(f1, 2) - f2)
		c4 := (2.07 - 0.23*math.Pow(f1, 2) - f2) * (0.63 - 0.07*math.Pow(f1, 2) - f2)
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2, c3, c4}
		return fitness
	}
	return cmop
}

func buildCecCmop7() types.CMOP {
	cmop := types.CMOP{
		ConstraintCount:   7,
		ObjectiveCount:    2,
		DecisionVariables: 6,
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrGreaterThanZero,
			types.EqualsOrLessThanZero,
			types.EqualsOrLessThanZero,
			types.EqualsOrLessThanZero,
			types.EqualsOrLessThanZero,
			types.EqualsOrLessThanZero,
			types.EqualsOrLessThanZero,
		},
		Name: "CEC2020-CMOP7",
		TrueParetoFrontFilename: "arraydata/cec2020/PF7.dat",
	}
	cmop.DecisionInterval = [][]float64{
		[]float64{0, 1},
		[]float64{78, 102},
		[]float64{33, 45},
		[]float64{27, 45},
		[]float64{27, 45},
		[]float64{27, 45},
	}
	cmop.Evaluate = func(x types.Genotype) types.Fitness {
		fitness := initFitness(cmop)
		g := 5.3578547*math.Pow(x[3], 2) + 0.8356891*x[1]*x[5] + 37.293239*x[1] - 10125.6023282166
		f1 := x[0]
		f2 := g * (1 - math.Sqrt(f1)/g)
		c1 := math.Pow(f1, 2) + math.Pow(f2, 2) - 1
		c2 := 85.334407 + 0.0056858*x[2]*x[5] + 0.0006262*x[1]*x[4] - 0.0022053*x[3]*x[5] - 92
		c3 := -85.334407 - 0.0056858*x[2]*x[5] - 0.0006262*x[1]*x[4] + 0.0022053*x[3]*x[5]
		c4 := 80.51249 + 0.0071317*x[2]*x[5] + 0.0029955*x[1]*x[2] + 0.0021813*math.Pow(x[3], 2) - 110
		c5 := -80.51249 - 0.0071317*x[2]*x[5] - 0.0029955*x[1]*x[2] - 0.0021813*math.Pow(x[3], 2) + 90
		c6 := 9.300961 + 0.0047026*x[3]*x[5] + 0.0012547*x[1]*x[3] + 0.0019085*x[3]*x[4] - 25
		c7 := -9.300961 - 0.0047026*x[3]*x[5] - 0.0012547*x[1]*x[3] - 0.0019085*x[3]*x[4] + 20
		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2, c3, c4, c5, c6, c7}
		return fitness
	}
	return cmop
}

func buildCecCmop8() types.CMOP {
	cmop := types.CMOP{
		ConstraintCount:   6,
		ObjectiveCount:    2,
		DecisionVariables: 8,
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrGreaterThanZero,
			types.EqualsOrGreaterThanZero,
			types.EqualsOrLessThanZero,
			types.EqualsOrLessThanZero,
			types.EqualsOrLessThanZero,
			types.EqualsOrLessThanZero,
		},
		Name: "CEC2020-CMOP8",
		TrueParetoFrontFilename: "arraydata/cec2020/PF8.dat",
	}
	cmop.DecisionInterval = [][]float64{
		[]float64{0, 1},
		[]float64{-10, 10},
		[]float64{-10, 10},
		[]float64{-10, 10},
		[]float64{-10, 10},
		[]float64{-10, 10},
		[]float64{-10, 10},
		[]float64{-10, 10},
	}

	cmop.Evaluate = func(x types.Genotype) types.Fitness {
		fitness := initFitness(cmop)
		g := math.Pow(x[1]-10, 2) + 5*math.Pow(x[2]-12, 2) + math.Pow(x[3], 4) + 3*math.Pow(x[4]-11, 2) + 10*math.Pow(x[5], 6) + 7*math.Pow(x[6], 2) + math.Pow(x[7], 4) - 4*x[6]*x[7] - 10*x[6] - 8*x[7] - 679.6300573745
		f1 := x[0]
		f2 := g * (1 - math.Sqrt(f1)/g)

		c1 := f1 + f2 - 1
		c2 := f1 + f2 - 1 - math.Abs(math.Sin(10*math.Pi*(f1-f2+1)))
		c3 := -127 + 2*math.Pow(x[1], 2) + 3*math.Pow(x[2], 4) + x[3] + 4*math.Pow(x[4], 2) + 5*x[5]
		c4 := -282 + 7*x[1] + 3*x[2] + 10*math.Pow(x[3], 2) + x[4] - x[5]
		c5 := -196 + 23*x[1] + math.Pow(x[2], 2) + 6*math.Pow(x[6], 2) - 8*x[7]
		c6 := 4*math.Pow(x[1], 2) + math.Pow(x[2], 2) - 3*x[1]*x[2] + 2*math.Pow(x[3], 2) + 5*x[6] - 11*x[7]

		fitness.ObjectiveValues = []float64{f1, f2}
		fitness.ConstraintValues = []float64{c1, c2, c3, c4, c5, c6}

		return fitness
	}

	return cmop
}

func CEC2020() cec2020 {
	return cec2020{
		CMOP1: buildCecCmop1(),
		CMOP2: buildCecCmop2(),
		CMOP3: buildCecCmop3(),
		CMOP4: buildCecCmop4(),
		CMOP5: buildCecCmop5(),
		CMOP6: buildCecCmop6(),
		CMOP7: buildCecCmop7(),
		CMOP8: buildCecCmop8(),
	}
}
