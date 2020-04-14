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
	/*
		for cr := 0; cr <= 10; cr++ {
			for f := 0; f <= 10; f++ {
				config.Moead.Cr = float64(cr) / 10
				config.Moead.F = float64(f) / 10
			}
		}
	*/
	simulator := simulator.NewSimulator(
		testsuites.LIR2D.SingleProblem(4),
		//testsuites.CEC2020.SingleProblem(0),
		10,
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
		CHM: types.R2S,
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
		CHM:                types.R2S,
		T:                  10,
		WeightDistribution: 101,
		Nr:                 2,
		Cr:                 0.7,
		F:                  0.8,
		DistributionIndex:  20.0,
	}

	exportconfig := configs.Export{
		ExportVideo:     true,
		PlotEval:        false,
		PrintGeneration: false,
		VideoMax:        3,
		VideoMin:        0,
	}

	sweeperconfig := configs.Sweeper{
		Sweep: true,
		Dir:   "results/sweep/",
		FR:    true,
		IGD:   true,
		HV:    true,
	}

	config := configs.Config{
		MaxFuncEvals: 100_000,

		Moead: moeadconfig,

		Stages: []types.StageType{types.Push /*types.BinarySearch,*/, types.Pull},

		Push:   pushconfig,
		Binary: binaryconfig,
		Pull:   pullconfig,

		CMOP: cmopconfig,

		CHM:           types.R2S,
		E:             econfig,
		IE:            ieconfig,
		R2S:           r2sconfig,
		Export:        exportconfig,
		HVCoefficient: 1.1,
		Sweeper:       sweeperconfig,
	}

	return config

}

```
