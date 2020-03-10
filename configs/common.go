package configs

import "github.com/CRAB-LAB-NTNU/PPS-BS/types"

type Config struct {

	// Common parameters
	MaxFuncEvals int

	//Optimizer
	Moead Moead

	//Stage specific
	Stages []types.StageType
	Push   Push
	Pull   Pull

	//CHM
	IE  ImprovedEpsilon
	E   Epsilon
	R2S R2S
}

type Moead struct {
	CHM                      types.CHMMethod
	T                        int
	WeightDistribution       int
	DecisionVariables        int
	Nr                       int
	Cr, F, DistributionIndex float64
}

type Epsilon struct {
	Cp float64
	Tc int
}
type ImprovedEpsilon struct {
	Tau, Alpha, Cp float64
	Tc             int
}
type R2S struct {
	FESc   int
	FESacd int
	Cs     int
	Val    float64
	ZMin   float64
}

type Push struct {
	Delta, Epsilon float64
	L              int
}

type Pull struct {
	CHM types.CHMMethod
}
