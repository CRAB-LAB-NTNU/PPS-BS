package cmops

import "github.com/CRAB-LAB-NTNU/PPS-BS/types"

var CEC2020 = types.TestSuite{
	Problems: []types.Cmop{
		cecCmop1,
		cecCmop2,
		cecCmop3,
		cecCmop4,
		cecCmop5,
		cecCmop6,
	},
	NumberOfProblems: 6,
	Name:             "CEC2020",
}

var LIR2D = types.TestSuite{
	Problems: []types.Cmop{
		lirCmop1,
		lirCmop2,
		lirCmop3,
		lirCmop4,
		lirCmop5,
		lirCmop6,
		lirCmop7,
		lirCmop8,
		lirCmop9,
		lirCmop10,
		lirCmop11,
		lirCmop12,
	},
	NumberOfProblems: 12,
	Name:             "Lir2D",
}

var LIR3D = types.TestSuite{
	Problems: []types.Cmop{
		lirCmop13,
		lirCmop14,
	},
	NumberOfProblems: 2,
	Name:             "Lir3D",
}
