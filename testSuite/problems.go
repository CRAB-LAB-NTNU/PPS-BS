package testSuite

import (
	"fmt"
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func CMOP1(x types.Genotype) types.Fitness {
	return types.Fitness{
		ObjectiveCount:  2,
		ConstraintCount: 2,
		ConstraintTypes: []types.ConstraintType{
			types.EqualsOrGreaterThanZero,
			types.EqualsOrGreaterThanZero,
		},
		ObjectiveValues: []float64{
			Objective1(x),
			Objective2(x),
		},
		ConstraintValues: []float64{
			Constraint1(x, Inner1),
			Constraint1(x, Inner2),
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
			Objective1(x),
			Objective3(x),
		},
		ConstraintValues: []float64{
			Constraint1(x, Inner1),
			Constraint1(x, Inner2),
		},
	}
}

func CMOP3(x types.Genotype) types.Fitness {
	fitness := CMOP1(x)
	fitness.ConstraintCount = 3
	fitness.ConstraintValues = append(fitness.ConstraintValues, Constraint2(x))
	return fitness
}

func CMOP4(x types.Genotype) types.Fitness {
	fitness := CMOP2(x)
	fitness.ConstraintCount = 3
	fitness.ConstraintValues = append(fitness.ConstraintValues, Constraint2(x))
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
			Objective4(x),
			Objective5(x),
		},
	}
	fitness.ConstraintValues = []float64{
		Constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		Constraint3(x, p, q, a, b, 1, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
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
			Objective4(x),
			Objective6(x),
		},
	}
	fitness.ConstraintValues = []float64{
		Constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		Constraint3(x, p, q, a, b, 1, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
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
		ObjectiveValues: []float64{Objective4(x), Objective5(x)},
	}
	fitness.ConstraintValues = []float64{
		Constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		Constraint3(x, p, q, a, b, 1, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		Constraint3(x, p, q, a, b, 2, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
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
		ObjectiveValues: []float64{Objective4(x), Objective6(x)},
	}
	fitness.ConstraintValues = []float64{
		Constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		Constraint3(x, p, q, a, b, 1, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		Constraint3(x, p, q, a, b, 2, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
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
		ObjectiveValues: []float64{Objective7(x), Objective8(x)},
	}
	fitness.ConstraintValues = []float64{
		Constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		Constraint4(x, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1], 2),
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
		ObjectiveValues: []float64{Objective7(x), Objective9(x)},
	}
	fitness.ConstraintValues = []float64{
		Constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		Constraint4(x, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1], 1),
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
		ObjectiveValues: []float64{Objective7(x), Objective9(x)},
	}
	fitness.ConstraintValues = []float64{
		Constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		Constraint4(x, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1], 2.1),
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
		ObjectiveValues: []float64{Objective7(x), Objective8(x)},
	}
	fitness.ConstraintValues = []float64{
		Constraint3(x, p, q, a, b, 0, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1]),
		Constraint4(x, fitness.ObjectiveValues[0], fitness.ObjectiveValues[1], 2.5),
	}
	return fitness
}

func CMOP13(x types.Genotype) types.Fitness {
	fitness := types.Fitness{
		ObjectiveCount: 3, ConstraintCount: 2,
		ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		ObjectiveValues: []float64{Objective10(x), Objective11(x), Objective12(x)},
	}
	g := math.Pow(fitness.ObjectiveValues[0], 2) + math.Pow(fitness.ObjectiveValues[1], 2) + math.Pow(fitness.ObjectiveValues[2], 2)
	fitness.ConstraintValues = []float64{
		Constraint5(x, g, 9, 4),
		Constraint5(x, g, 3.61, 3.24),
	}
	return fitness
}

func CMOP14(x types.Genotype) types.Fitness {
	fitness := types.Fitness{
		ObjectiveCount: 3, ConstraintCount: 3,
		ConstraintTypes: []types.ConstraintType{types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero, types.EqualsOrGreaterThanZero},
		ObjectiveValues: []float64{Objective10(x), Objective11(x), Objective12(x)},
	}
	g := math.Pow(fitness.ObjectiveValues[0], 2) + math.Pow(fitness.ObjectiveValues[1], 2) + math.Pow(fitness.ObjectiveValues[2], 2)
	fitness.ConstraintValues = []float64{
		Constraint5(x, g, 9, 4),
		Constraint5(x, g, 3.61, 3.24),
		Constraint5(x, g, 3.0625, 2.56),
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
