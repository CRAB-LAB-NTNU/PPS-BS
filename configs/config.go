package configs

import "github.com/CRAB-LAB-NTNU/PPS-BS/types"

type Config struct {

	// Common parameters
	MaxFuncEvals int
	Export       Export

	//Optimizer
	Moead Moead

	//Stage specific
	Stages []types.StageType
	Push   Push
	Binary Binary
	Pull   Pull

	//CHM
	CHM types.CHMMethod
	IE  ImprovedEpsilon
	E   Epsilon
	R2S R2S
}
