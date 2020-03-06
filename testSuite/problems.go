package testSuite

import (
	"fmt"
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

type LIRCMOP struct {
	numberOfObjectives, numberOfConstraints int
	name                                    string
	CalcFunc                                func(types.Genotype) types.Fitness
}

func (lircmop *LIRCMOP) Initialize(problem int) {
	lircmop.name = fmt.Sprintf("%s%d", "LIRCMOP", problem)
	switch problem {
	case 1:
		lircmop.numberOfObjectives = 2
		lircmop.numberOfConstraints = 2
		lircmop.CalcFunc = CMOP1
	case 2:
		lircmop.numberOfObjectives = 2
		lircmop.numberOfConstraints = 2
		lircmop.CalcFunc = CMOP2
	case 3:
		lircmop.numberOfObjectives = 2
		lircmop.numberOfConstraints = 3
		lircmop.CalcFunc = CMOP3
	case 4:
		lircmop.numberOfObjectives = 2
		lircmop.numberOfConstraints = 3
		lircmop.CalcFunc = CMOP4

	case 5:
		lircmop.numberOfObjectives = 2
		lircmop.numberOfConstraints = 2
		lircmop.CalcFunc = CMOP5
	case 6:
		lircmop.numberOfObjectives = 2
		lircmop.numberOfConstraints = 2
		lircmop.CalcFunc = CMOP6
	case 7:
		lircmop.numberOfObjectives = 2
		lircmop.numberOfConstraints = 3
		lircmop.CalcFunc = CMOP7
	case 8:
		lircmop.numberOfObjectives = 2
		lircmop.numberOfConstraints = 3
		lircmop.CalcFunc = CMOP8
	case 9:
		lircmop.numberOfObjectives = 2
		lircmop.numberOfConstraints = 2
		lircmop.CalcFunc = CMOP9
	case 10:
		lircmop.numberOfObjectives = 2
		lircmop.numberOfConstraints = 2
		lircmop.CalcFunc = CMOP10
	case 11:
		lircmop.numberOfObjectives = 2
		lircmop.numberOfConstraints = 2
		lircmop.CalcFunc = CMOP11
	case 12:
		lircmop.numberOfObjectives = 2
		lircmop.numberOfConstraints = 2
		lircmop.CalcFunc = CMOP12
	case 13:
		lircmop.numberOfObjectives = 3
		lircmop.numberOfConstraints = 2
		lircmop.CalcFunc = CMOP13
	case 14:
		lircmop.numberOfObjectives = 3
		lircmop.numberOfConstraints = 3
		lircmop.CalcFunc = CMOP14

	}
}

func (lircmop LIRCMOP) Calculate(x types.Genotype) types.Fitness {
	return lircmop.CalcFunc(x)
}

func (lircmop LIRCMOP) Name() string {
	return lircmop.name
}

func (lircmop LIRCMOP) NumberOfConstraints() int {
	return lircmop.numberOfConstraints
}

func (lircmop LIRCMOP) NumberOfObjectives() int {
	return lircmop.numberOfObjectives
}

func CMOP1(x types.Genotype) types.Fitness {
	return types.Fitness{
		ObjectiveCount:  2,
		ConstraintCount: 2,
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrGreaterThanZero,
			types.EqualsOrGreaterThanZero,
		},
		ObjectiveValues: []float64{
			objective1(x),
			objective2(x),
		},
		ConstraintValues: []float64{
			constraint1(x, inner1),
			constraint1(x, inner2),
		},
	}
}

func CMOP2(x types.Genotype) types.Fitness {
	return types.Fitness{
		ObjectiveCount:  2,
		ConstraintCount: 2,
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrGreaterThanZero,
			types.EqualsOrGreaterThanZero,
		},
		ObjectiveValues: []float64{
			objective1(x),
			objective3(x),
		},
		ConstraintValues: []float64{
			constraint1(x, inner1),
			constraint1(x, inner2),
		},
	}
}

func CMOP3(x types.Genotype) types.Fitness {
	fitness := CMOP1(x)
	fitness.ConstraintCount = 3
	fitness.ConstraintValues = append(fitness.ConstraintValues, constraint2(x))
	return fitness
}

func CMOP4(x types.Genotype) types.Fitness {
	fitness := CMOP2(x)
	fitness.ConstraintCount = 3
	fitness.ConstraintValues = append(fitness.ConstraintValues, constraint2(x))
	return fitness
}

func CMOP5(x types.Genotype) types.Fitness {
	p := []float64{1.6, 2.5}
	q := []float64{1.6, 2.5}
	a := []float64{2, 2}
	b := []float64{4, 8}
	fitness := types.Fitness{
		ObjectiveCount:  2,
		ConstraintCount: 2,
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrGreaterThanZero,
			types.EqualsOrGreaterThanZero,
		},
		ObjectiveValues: []float64{
			objective4(x),
			objective5(x),
		},
	}
	fitness.ConstraintValues = []float64{
		constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		constraint3(x, p, q, a, b, 1, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
	}
	return fitness
}

func CMOP6(x types.Genotype) types.Fitness {
	p := []float64{1.8, 2.8}
	q := []float64{1.8, 2.8}
	a := []float64{2, 2}
	b := []float64{8, 8}
	fitness := types.Fitness{
		ObjectiveCount:  2,
		ConstraintCount: 2,
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrGreaterThanZero,
			types.EqualsOrGreaterThanZero,
		},
		ObjectiveValues: []float64{
			objective4(x),
			objective6(x),
		},
	}
	fitness.ConstraintValues = []float64{
		constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		constraint3(x, p, q, a, b, 1, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
	}
	return fitness
}

func CMOP7(x types.Genotype) types.Fitness {
	p := []float64{1.2, 2.25, 3.5}
	q := []float64{1.2, 2.25, 3.5}
	a := []float64{2, 2.5, 2.5}
	b := []float64{6, 12, 10}
	fitness := types.Fitness{
		ObjectiveCount: 2, ConstraintCount: 3,
		ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		ObjectiveValues: []float64{objective4(x), objective5(x)},
	}
	fitness.ConstraintValues = []float64{
		constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		constraint3(x, p, q, a, b, 1, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		constraint3(x, p, q, a, b, 2, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
	}
	return fitness
}

func CMOP8(x types.Genotype) types.Fitness {
	p := []float64{1.2, 2.25, 3.5}
	q := []float64{1.2, 2.25, 3.5}
	a := []float64{2, 2.5, 2.5}
	b := []float64{6, 12, 10}
	fitness := types.Fitness{
		ObjectiveCount: 2, ConstraintCount: 3,
		ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		ObjectiveValues: []float64{objective4(x), objective6(x)},
	}
	fitness.ConstraintValues = []float64{
		constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		constraint3(x, p, q, a, b, 1, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		constraint3(x, p, q, a, b, 2, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
	}
	return fitness
}

func CMOP9(x types.Genotype) types.Fitness {
	p := []float64{1.4}
	q := []float64{1.4}
	a := []float64{1.5}
	b := []float64{6.0}
	fitness := types.Fitness{
		ObjectiveCount: 2, ConstraintCount: 2,
		ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		ObjectiveValues: []float64{objective7(x), objective8(x)},
	}
	fitness.ConstraintValues = []float64{
		constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		constraint4(x, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1], 2),
	}
	return fitness
}

func CMOP10(x types.Genotype) types.Fitness {
	p := []float64{1.1}
	q := []float64{1.2}
	a := []float64{2.0}
	b := []float64{4.0}
	fitness := types.Fitness{
		ObjectiveCount: 2, ConstraintCount: 2,
		ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		ObjectiveValues: []float64{objective7(x), objective9(x)},
	}
	fitness.ConstraintValues = []float64{
		constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		constraint4(x, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1], 1),
	}
	return fitness
}

func CMOP11(x types.Genotype) types.Fitness {
	p := []float64{1.2}
	q := []float64{1.2}
	a := []float64{1.5}
	b := []float64{5.0}
	fitness := types.Fitness{
		ObjectiveCount: 2, ConstraintCount: 2,
		ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		ObjectiveValues: []float64{objective7(x), objective9(x)},
	}
	fitness.ConstraintValues = []float64{
		constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		constraint4(x, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1], 2.1),
	}
	return fitness
}

func CMOP12(x types.Genotype) types.Fitness {
	p := []float64{1.6}
	q := []float64{1.6}
	a := []float64{1.5}
	b := []float64{6.0}
	fitness := types.Fitness{
		ObjectiveCount: 2, ConstraintCount: 2,
		ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		ObjectiveValues: []float64{objective7(x), objective8(x)},
	}
	fitness.ConstraintValues = []float64{
		constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		constraint4(x, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1], 2.5),
	}
	return fitness
}

func CMOP13(x types.Genotype) types.Fitness {
	fitness := types.Fitness{
		ObjectiveCount: 3, ConstraintCount: 2,
		ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		ObjectiveValues: []float64{objective10(x), objective11(x), objective12(x)},
	}
	g := math.Pow(fitness.ObjectiveValues[0], 2) + math.Pow(fitness.ObjectiveValues[1], 2) + math.Pow(fitness.ObjectiveValues[2], 2)
	fitness.ConstraintValues = []float64{
		constraint5(x, g, 9, 4),
		constraint5(x, g, 3.61, 3.24),
	}
	return fitness
}

func CMOP14(x types.Genotype) types.Fitness {
	fitness := types.Fitness{
		ObjectiveCount: 3, ConstraintCount: 3,
		ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		ObjectiveValues: []float64{objective10(x), objective11(x), objective12(x)},
	}
	g := math.Pow(fitness.ObjectiveValues[0], 2) + math.Pow(fitness.ObjectiveValues[1], 2) + math.Pow(fitness.ObjectiveValues[2], 2)
	fitness.ConstraintValues = []float64{
		constraint5(x, g, 9, 4),
		constraint5(x, g, 3.61, 3.24),
		constraint5(x, g, 3.0625, 2.56),
	}
	return fitness
}

func RobotGripper(x types.Genotype) types.Fitness {
	fitness := types.Fitness{
		ObjectiveCount: 2, ConstraintCount: 7,
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrGreaterThanZero,
			types.EqualsOrGreaterThanZero,
			types.EqualsOrGreaterThanZero,
			types.EqualsOrGreaterThanZero,
			types.EqualsOrGreaterThanZero,
			types.EqualsOrGreaterThanZero,
			types.EqualsOrGreaterThanZero,
		},
	}
	converted := convertGenotype(x)
	fuzzyTest(converted)
	fmt.Println(converted)
	yMin, yG, yMax, zMax, P := 50.0, 150.0, 100.0, 100.0, 100.0

	fitness.ObjectiveValues = []float64{
		robotObjective1(converted, zMax, P),
		robotObjective2(converted),
	}

	fitness.ConstraintValues = []float64{
		robotConstraint1(converted, yMin, zMax),
		robotConstraint2(converted, zMax),
		robotConstraint3(converted, yMax),
		robotConstraint4(converted, yG),
		robotConstraint5(converted),
		robotConstraint6(converted, zMax),
		robotConstraint7(converted, zMax),
	}

	return fitness
}
