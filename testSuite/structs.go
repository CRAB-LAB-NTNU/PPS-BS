package testSuite

type Fitness struct {
	Objectives, SoftConstraints map[string]float64
	HardConstraints             map[string]bool
}
