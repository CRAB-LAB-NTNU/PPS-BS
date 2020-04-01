package simulator

import (
	"fmt"

	"github.com/CRAB-LAB-NTNU/PPS-BS/chm"
	"github.com/CRAB-LAB-NTNU/PPS-BS/configs"
	"github.com/CRAB-LAB-NTNU/PPS-BS/metrics"
	"github.com/CRAB-LAB-NTNU/PPS-BS/optimisers"
	"github.com/CRAB-LAB-NTNU/PPS-BS/pps"
	"github.com/CRAB-LAB-NTNU/PPS-BS/stages"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

type Simulator struct {
	TestSuite types.TestSuite
	Runs      int
	Config    configs.Config
	results   []metrics.Results
}

func NewSimulator(testSuite types.TestSuite, runs int, config configs.Config) Simulator {
	s := Simulator{
		TestSuite: testSuite,
		Runs:      runs,
		Config:    config,
	}

	return s
}

func (s *Simulator) Simulate() {
	for _, cmop := range s.TestSuite.Problems {
		type indChan chan []types.Individual
		channel := make(indChan)

		// Start (s.Runs) new Go routines for and send pps result to channel.
		for r := 0; r < s.Runs; r++ {
			go func(cmop types.Cmop, r int, channel indChan) {
				pps := s.setupInstance(cmop)
				channel <- pps.Run()
			}(cmop, r, channel)
		}

		// Create result struct, then move all items from channel to result.
		result := metrics.Results{
			ParetoFront:          cmop.TrueParetoFront(),
			HyperVolumeReference: metrics.HVReferenceNadirTimes(s.Config.HVCoefficient, cmop),
		}
		for r := 0; r < s.Runs; r++ {
			result.Add(<-channel)
		}

		s.results = append(s.results, result)
	}
	s.printSweep()
}

func (s *Simulator) printSweep() {
	for i, r := range s.results {
		fmt.Println(s.TestSuite.Problems[i].Name, s.Config.Moead.Cr, s.Config.Moead.F, r.FeasibilityRate(), r.IGD.Mean(), r.HV.Mean())
	}
}

/*
func (s *Simulator) printResults(p int, cmopName string, runTime time.Duration) {
	fmt.Println("PROBLEM:", cmopName)
	fmt.Println("Run time:", runTime)
		fmt.Println("Stages:", pps.Stages())
		fmt.Println("Constraint Handling Method:", pps.MOEA().CHM().Name())

		fmt.Println("BEST:", arrays.Min(s.results[p].Values()...))
		fmt.Println("MEAN:", s.results[p].Mean())
		fmt.Println("VAR:", s.results[p].Variance())
		fmt.Println("STD:", s.results[p].StandardDeviation())
		fmt.Println()
	}
*/

func (s *Simulator) setupInstance(cmop types.Cmop) pps.PPS {

	var stages []types.Stage

	for i := range s.Config.Stages {
		stages = append(stages, s.setupStage(i, cmop.ObjectiveCount))
	}

	chm := s.setupChm(cmop.ConstraintCount)

	moea := s.setupMoea(cmop, chm)

	pps := s.setupPps(cmop, moea, stages)

	return pps

}

func (s *Simulator) setupStage(pos int, numberOfObjectives int) types.Stage {
	switch s.Config.Stages[pos] {
	case types.Push:
		return stages.NewPush(
			s.Config.Push.Delta,
			s.Config.Push.Epsilon,
			s.Config.Push.L,
			s.Config.MaxFuncEvals,
			numberOfObjectives)
	case types.BinarySearch:
		return stages.NewBinary(
			s.Config.Binary.MinDistance,
			s.Config.Binary.Fcp)
	case types.Pull:
		return stages.NewPull()
	}
	panic("Error setting up stage")
}

func (s *Simulator) setupChm(numberOfConstraints int) types.CHM {
	switch s.Config.CHM {
	case types.Epsilon:
		return chm.NewE(
			s.Config.E.Cp,
			s.Config.E.Tc,
			s.Config.MaxFuncEvals)
	case types.ImprovedEpsilon:
		return chm.NewIE(s.Config.IE.Tau,
			s.Config.IE.Alpha,
			s.Config.IE.Cp,
			s.Config.IE.Tc,
			s.Config.MaxFuncEvals)
	case types.R2S:
		return chm.NewR2S(s.Config.R2S.FESc,
			s.Config.R2S.FESacd,
			s.Config.R2S.Cs,
			s.Config.R2S.Val,
			s.Config.R2S.ZMin,
			numberOfConstraints,
			s.Config.MaxFuncEvals)
	}
	panic("Error setting up CHM")
}

func (s *Simulator) setupMoea(cmop types.Cmop, chm types.CHM) types.MOEA {
	return optimisers.NewMoead(
		cmop, chm,
		s.Config.Moead.T,
		s.Config.Moead.WeightDistribution,
		s.Config.Moead.DecisionVariables,
		s.Config.Moead.Nr,
		s.Config.Moead.F,
		s.Config.Moead.Cr,
		s.Config.Moead.DistributionIndex,
		s.Config.MaxFuncEvals,
	)
}

func (s *Simulator) setupPps(cmop types.Cmop, moea types.MOEA, stages []types.Stage) pps.PPS {
	return pps.NewPPS(
		cmop,
		moea,
		stages,
		s.Config.Export)
}
