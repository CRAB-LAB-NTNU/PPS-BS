package configs

import "github.com/CRAB-LAB-NTNU/PPS-BS/types"

type Moead struct {
	CHM                      types.CHMMethod
	T                        int
	WeightDistribution       int
	DecisionVariables        int
	Nr                       int
	Cr, F, DistributionIndex float64
}
