package configs

import "github.com/CRAB-LAB-NTNU/PPS-BS/types"

type PPS struct {
	ExportVideo        bool
	Runs               int
	Metric             func([]types.Individual, [][]float64) float64
	VideoMax, VideoMin float64
}
