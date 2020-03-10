package types

import (
	"math"
)

// CMOP is a interface describing a multi objective optimisation problem
type CMOP interface {
	NumberOfObjectives() int
	NumberOfConstraints() int
	Name() string
	Calculate(Genotype) Fitness
}

// ConstraintType describes which type of constraint it is. Either a equals-or-less-than or equals-or-greater-than constraint.
// Should we assume only inequality constraints by using a small delta? Seems like most approaches do
type ConstraintType int

const (
	//EqualsOrLessThanZero the constraint is a less-than-or-equals-zero constraint
	EqualsOrLessThanZero ConstraintType = iota + 1
	//EqualsOrGreaterThanZero the constraint is a less-than-or-greater-zero constraint
	EqualsOrGreaterThanZero ConstraintType = iota + 1
)

// MOEA is an interface describing Multi Objective Evolutionary Algorithms
type MOEA interface {
	MaxFuncEvals() int
	MaxViolation() float64
	Population() []Individual
	Initialise()
	InitialiseCHM()
	FunctionEvaluations() int
	FeasibleRatio() float64
	Reset()
	Ideal() []float64
	Archive() []Individual
	Evolve(StageType)
	ResetBinary()
	IsBinarySearch() bool
	BinaryDone() bool
	GetCHM() CHM
}

// Individual is an interface describing an individual in a population
type Individual interface {
	Genotype() Genotype //TODO: se på måter å gjøre dette mer generelt senere
	Fitness() Fitness
	UpdateFitness(CMOP) Fitness
	Copy() Individual
}

type Fitness struct {
	ObjectiveCount, ConstraintCount   int
	ObjectiveValues, ConstraintValues []float64
	ConstraintTypes                   []ConstraintType
}

//Violation returns the violation of a constraint at position c.
func (f Fitness) Violation(c int) float64 {
	if f.ConstraintTypes[c] == EqualsOrGreaterThanZero && f.ConstraintValues[c] < 0 {
		return math.Abs(f.ConstraintValues[c])
	} else if f.ConstraintTypes[c] == EqualsOrLessThanZero && f.ConstraintValues[c] > 0 {
		return f.ConstraintValues[c]
	}
	return 0.0
}

// TotalViolation returns the total constraint violation of all constraints.
func (f Fitness) TotalViolation() float64 {
	var total float64
	for pos, constraintVal := range f.ConstraintValues {
		if f.ConstraintTypes[pos] == EqualsOrGreaterThanZero && constraintVal < 0 {
			total += math.Abs(constraintVal)
		} else if f.ConstraintTypes[pos] == EqualsOrLessThanZero && constraintVal > 0 {
			total += constraintVal
		}
	}
	return total
}

type Genotype []float64
