package types

// CMOP is a interface describing a multi objective optimisation problem
type CMOP struct {
	NumberOfObjectives int
	Calculate          func(Genotype) Fitness
}

/*
// Objective describes an objective of a problem
type Objective struct {
	Type     ObjectiveType
	Function ObjectiveFunction
}
*/

// ObjectiveType says if the objective is a minimisation or maximisation problem
type ObjectiveType int

const (
	// Minimisation the objective is a minimisation problem
	Minimisation ObjectiveType = iota + 1
	// Maximisation the objective is a maximisation problem
	Maximisation
)

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
	MaxGeneration() int
	MaxViolation() float64
	Population() []Individual
	Initialise()
	Evolve(Stage, []float64)
	Crossover([]Individual) []Individual
}

// Individual is an interface describing an individual in a population
type Individual interface {
	Genotype() Genotype //TODO: se på måter å gjøre dette mer generelt senere
	Fitness() Fitness
	UpdateFitness(CMOP) Fitness
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
	ObjectiveTypes                    []ObjectiveType
	ConstraintTypes                   []ConstraintType
}

type Genotype []float64
