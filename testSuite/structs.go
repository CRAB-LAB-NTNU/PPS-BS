package testSuite

type Fitness struct {
	Objectives, SoftConstraints map[string]float64
	InequalityConstraints       map[string]bool
}
