# PPS-BS

Example main class setting up all the config for the framework:

```
package main

import (
	"math/rand"
	"time"

	"github.com/CRAB-LAB-NTNU/PPS-BS/configs"
	"github.com/CRAB-LAB-NTNU/PPS-BS/metrics"
	"github.com/CRAB-LAB-NTNU/PPS-BS/simulator"
	"github.com/CRAB-LAB-NTNU/PPS-BS/testsuites"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	config := setupConfigs()

	simulator := simulator.NewSimulator(
		testsuites.LIR2D,
		30,
		config)

	simulator.Simulate()

}

func setupConfigs() configs.Config {
	// Stages
	pushconfig := configs.Push{
		Delta:   1.0 / 1000000.0,
		Epsilon: 1.0 / 1000.0,
		L:       20,
	}

	binaryconfig := configs.Binary{
		MinDistance: 0.05,
		Fcp:         0.5,
	}

	pullconfig := configs.Pull{
		CHM: types.ImprovedEpsilon,
	}

	//CMOP
	cmopconfig := configs.CMOP{
		Problem: 1,
	}

	//CHMs

	r2sconfig := configs.R2S{
		FESc:   300000 * 9 / 10,
		FESacd: 100000,
		Cs:     25,
		Val:    0.01,
		ZMin:   3,
	}

	ieconfig := configs.ImprovedEpsilon{
		Tau:   0.1,
		Alpha: 0.95,
		Cp:    2.0,
		Tc:    800,
	}

	econfig := configs.Epsilon{
		Cp: 2.0,
		Tc: 800,
	}

	//Optimizers

	moeadconfig := configs.Moead{
		CHM:                types.ImprovedEpsilon,
		T:                  30,
		WeightDistribution: 301,
		DecisionVariables:  30,
		Nr:                 2,
		Cr:                 1,
		F:                  0.5,
		DistributionIndex:  20.0,
	}

	exportconfig := configs.Export{
		ExportVideo:     false,
		PlotEval:        false,
		PrintGeneration: false,
		VideoMax:        3,
		VideoMin:        0,
		Metric:          metrics.InvertedGenerationalDistance,
	}

	config := configs.Config{
		MaxFuncEvals: 300_000,

		Moead: moeadconfig,

		Stages: []types.StageType{types.Push, types.Pull},

		Push:   pushconfig,
		Binary: binaryconfig,
		Pull:   pullconfig,

		CMOP: cmopconfig,

		CHM:    types.ImprovedEpsilon,
		E:      econfig,
		IE:     ieconfig,
		R2S:    r2sconfig,
		Export: exportconfig,
	}

	return config

}
```