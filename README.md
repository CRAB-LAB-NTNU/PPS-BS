# PPS-BS

Example main class setting up all the config for the framework:

```
package main

import (
	"math/rand"
	"time"

	"github.com/CRAB-LAB-NTNU/PPS-BS/configs"
	"github.com/CRAB-LAB-NTNU/PPS-BS/simulator"
	"github.com/CRAB-LAB-NTNU/PPS-BS/testsuites"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	config := setupConfigs()
	sim := simulator.NewSimulator(testsuites.LIR2D, 4, config)
	sim.Simulate()

	config.Moead.WeightDistribution = 101
	config.MaxFuncEvals = 100_000
	config.Moead.Cr = 0.9
	config.Moead.F = 1.0
	sim = simulator.NewSimulator(testsuites.CEC2020, 4, config)
	sim.Simulate()
}

func setupConfigs() configs.Config {
	// Stages
	pushconfig := configs.Push{
		Delta:   1.0 / 1000000.0,
		Epsilon: 1.0 / 1000.0,
		L:       20,
	}

	binaryconfig := configs.Binary{
		MinDistance: 0.03,
		Fcp:         0.1,
	}

	pullconfig := configs.Pull{
		CHM: types.Epsilon,
	}

	//CMOP
	cmopconfig := configs.CMOP{
		Problem: 1,
	}

	//CHMs

	r2sconfig := configs.R2S{
		FESc:   300000.0 * 9.0 / 10.0,
		NUMacd: 10,
		Val:    0.01,
		ZMin:   3.0,
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
		CHM:                types.Epsilon,
		T:                  30,
		WeightDistribution: 301,
		Nr:                 2,
		Cr:                 1.0,
		F:                  0.5,
		DistributionIndex:  20.0,
	}

	exportconfig := configs.Export{
		ExportVideo:     false,
		PlotEval:        false,
		PrintGeneration: false,
		VideoMax:        3,
		VideoMin:        0,
	}

	sweeperconfig := configs.Sweeper{
		Sweep:      true,
		Dir:        "results/sweep/",
		FR:         true,
		IGD:        true,
		HV:         true,
		ArchiveIGD: true,
		ArchiveHV:  true,
		Phase:      true,
	}

	config := configs.Config{
		MaxFuncEvals: 300_000,

		Moead: moeadconfig,

		Stages: []types.StageType{types.Push, types.BinarySearch, types.Pull},

		Push:   pushconfig,
		Binary: binaryconfig,
		Pull:   pullconfig,

		CMOP: cmopconfig,

		CHM:           types.Epsilon,
		E:             econfig,
		IE:            ieconfig,
		R2S:           r2sconfig,
		Export:        exportconfig,
		HVCoefficient: 1.1,
		Sweeper:       sweeperconfig,
		TimeStamp:     time.Now().Format(time.Stamp),
	}

	return config

}

```
