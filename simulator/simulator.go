package simulator

import (
	"fmt"

	"github.com/CRAB-LAB-NTNU/PPS-BS/chm"
	"github.com/CRAB-LAB-NTNU/PPS-BS/configs"
	"github.com/CRAB-LAB-NTNU/PPS-BS/optimisers"
	"github.com/CRAB-LAB-NTNU/PPS-BS/pps"
	"github.com/CRAB-LAB-NTNU/PPS-BS/stages"
	"github.com/CRAB-LAB-NTNU/PPS-BS/testSuite"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

type Simulator struct {
	Problems []int
	Runs     int
	Config   configs.Config
	results  []types.Results
}

func NewSimulator(problems []int, runs int, config configs.Config) Simulator {
	s := Simulator{
		Problems: problems,
		Runs:     runs,
		Config:   config,
	}

	for i := 0; i <= len(problems); i++ {
		s.results = append(s.results, types.Results{})
	}

	return s
}
func (s *Simulator) Simulate() {
	for _, p := range s.Problems {
		cmop := testSuite.NewLIRCMOP(p)

		fmt.Println("Starting run of", cmop.Name())
		var pps pps.PPS
		for r := 0; r < s.Runs; r++ {
			pps = s.setupInstance(cmop)
			s.results[p].Add(pps.Run())
		}

		s.printResults(p, pps)

	}
}

func (s *Simulator) printResults(p int, pps pps.PPS) {
	fmt.Println("PROBLEM:", pps.CMOP().Name())
	fmt.Println("Stages:", pps.Stages())
	fmt.Println("Constraint Handling Method:", pps.MOEA().CHM().Name())

	fmt.Println("MEAN:", s.results[p].Mean())
	fmt.Println("VAR:", s.results[p].Variance())
	fmt.Println("STD:", s.results[p].StandardDeviation())
	fmt.Println()
}

func (s *Simulator) setupInstance(cmop types.CMOP) pps.PPS {

	var stages []types.Stage

	for i := range s.Config.Stages {
		stages = append(stages, s.setupStage(i, cmop.NumberOfObjectives()))
	}

	chm := s.setupChm(cmop.NumberOfConstraints())

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

func (s *Simulator) setupMoea(cmop types.CMOP, chm types.CHM) types.MOEA {
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

func (s *Simulator) setupPps(cmop types.CMOP, moea types.MOEA, stages []types.Stage) pps.PPS {
	return pps.NewPPS(
		cmop,
		moea,
		stages,
		s.Config.Export)
}
