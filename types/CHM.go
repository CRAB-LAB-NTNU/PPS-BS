package types

type CHM interface {
	Name() string
	Threshold(int) float64
	Initialise()
	Update(int, float64)
	Violation(int, Fitness) float64
}

type CHMMethod int

const (
	//Epsilon method uses normal epsilon
	Epsilon CHMMethod = iota + 1
	//ImprovedEpsilon uses the improved epsilon constraint handling method
	ImprovedEpsilon
	//R2S uses the Reduces Search Space method to handle constraints
	R2S
)
