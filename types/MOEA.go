package types

import (
	"math"
)

/*
// Objective describes an objective of a problem
type Objective struct {
	Type     ObjectiveType
	Function ObjectiveFunction
}
*/

/*
// ObjectiveFunction describes the function for maximisation or minimisation for an objective
type ObjectiveFunction func(Genotype) float64
*/

/*
// Constraint describes a constraint for a problem
type Constraint interface {
	Type() ConstraintType
	Function() float64 //TODO: Do we have a better name?
}
*/
// ConstraintType describes which type of constraint it is. Either a equals-or-less-than or equals-or-greater-than constraint.
// Should we assume only inequality constraints by using a small delta? Seems like most approaches do
type ConstraintType int

const (
	//EqualsOrLessThanZero the constraint is a less-than-or-equals-zero constraint
	EqualsOrLessThanZero ConstraintType = iota + 1
	//EqualsOrGreaterThanZero the constraint is a less-than-or-greater-zero constraint
	EqualsOrGreaterThanZero ConstraintType = iota + 1
)

/*
// ConstraintFunction evaluates the constraint violation of an individual
type ConstraintFunction func(Genotype) float64
*/
// MOEA is an interface describing Multi Objective Evolutionary Algorithms
type MOEA interface {
	MaxEvaluations() int
	MaxViolation() float64
	Population() []Individual
	Initialise(CMOP)
	FunctionEvaluations() int
	FeasibleRatio() float64
	Reset()
	Ideal() []float64
	Archive() []Individual
	Evolve(Stage, []float64)
	ResetBinary()
	IsBinarySearch() bool
	BinaryDone() bool
}

// Individual is an interface describing an individual in a population
type Individual interface {
	Genotype() Genotype //TODO: se på måter å gjøre dette mer generelt senere
	Fitness() Fitness
	UpdateFitness() Fitness
	Copy() Individual
	Initialise()
	//Mutate()
	//ConstraintViolation()
	//UpdateConstraintViolation()
}

/*
type Fitness struct {
	Objectives, SoftConstraints map[string]float64
	InequalityConstraints       map[string]bool
}
*/
type Fitness struct {
	ObjectiveCount, ConstraintCount   int
	ObjectiveValues, ConstraintValues []float64
	ConstraintTypes                   []ConstraintType
}

/*
ConstraintViolation returns the total violation of an individual.
Evaluates both Greater or Less than zero.
*/
func (f Fitness) ConstraintViolation() float64 {
	var total float64
	for i := range f.ConstraintTypes {
		conVal := f.ConstraintValues[i]
		switch f.ConstraintTypes[i] {
		case EqualsOrGreaterThanZero:
			total += math.Abs(math.Min(conVal, 0))
		case EqualsOrLessThanZero:
			total += math.Max(conVal, 0)
		}
	}
	return total
}

/*
Feasible returns true if an individual is feasible or false if it's infeasible,
according to it's constraint values.
*/
func (f Fitness) Feasible() bool {
	return f.ConstraintViolation() <= 0
}

type Genotype []float64
