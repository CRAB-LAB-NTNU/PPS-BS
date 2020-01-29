package pps

import "github.com/CRAB-LAB-NTNU/PPS-BS/testSuite"

// PPS is a struct describing the contents of the Push & Pull Framework
type PPS struct {
	Cmop                     CMOP
	Moea                     MOEA
	Stage                    Stage
	IdealPoints, NadirPoints [][]float64
	RK, Delta, Epsilon       float64
	TC, L                    int
}

// Stage is an enum defining which phase PPS is in. Can either be Push or Pull
type Stage int

const (
	// Push PPS is in the push stage and constraints are ignored
	Push Stage = iota + 1
	// Pull PPS is in the pull stage and constraints are handled
	Pull
)

// CMOP is a interface describing a multi objective optimisation problem
type CMOP interface {
	NumberOfObjectives() int
	Objectives() []Objective
	Constraints() []Constraint
}

// Objective describes an objective of a problem
type Objective struct {
	Type     ObjectiveType
	Function ObjectiveFunction
}

// ObjectiveType says if the objective is a minimisation or maximisation problem
type ObjectiveType int

const (
	// Minimisation the objective is a minimisation problem
	Minimisation ObjectiveType = iota + 1
	// Maximisation the objective is a maximisation problem
	Maximisation
)

// ObjectiveFunction describes the function for maximisation or minimisation for an objective
type ObjectiveFunction func(Individual) float64

// Constraint describes a constraint for a problem
type Constraint struct {
	Type     ConstraintType
	Function ConstraintFunction //TODO: Do we have a better name?
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

// ConstraintFunction evaluates the constraint violation of an individual
type ConstraintFunction func(Individual) float64

// MOEA is an interface describing Multi Objective Evolutionary Algorithms
type MOEA interface {
	Initialise()
	MaxGeneration() int
	Population() []Individual
	InitialisePopulation() []Individual
	Evaluate() testSuite.Fitness
	Evolve(Stage) []Individual
	Cossover() []Individual
	Mutate() Individual
}

// Individual is an interface describing an individual in a population
type Individual interface {
	Genotype()
	Fitness()
	UpdateFitness()
	Mutate()
	ConstraintViolation()
	UpdateConstraintViolation()
}
