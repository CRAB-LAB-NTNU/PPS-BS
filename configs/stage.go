package configs

import "github.com/CRAB-LAB-NTNU/PPS-BS/types"

type Push struct {
	Delta, Epsilon float64
	L              int
}

type Binary struct {
	MinDistance float64
	Fcp         float64
}
type Pull struct {
	CHM types.CHMMethod
}
