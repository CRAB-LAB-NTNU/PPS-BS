package testsuites

import "github.com/CRAB-LAB-NTNU/PPS-BS/types"

// CEC2020 problems from http://www.escience.cn/people/yongwang1/competition2020.html
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
	Name:             "MW",
}

// LIR2D contains the LIR-CMOP1-12 two objective problems.
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
	Name:             "LIR2D",
}

// LIR3D contains the LIR-CMOP13-14 three objective problems.
var LIR3D = types.TestSuite{
	Problems: []types.Cmop{
		lirCmop13,
		lirCmop14,
	},
	NumberOfProblems: 2,
	Name:             "LIR3D",
}
