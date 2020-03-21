package types

import (
	"log"

	"github.com/CRAB-LAB-NTNU/PPS-BS/filehandling"
)

// Cmop is a interface describing a multi objective optimisation problem
type Cmop struct {
	ObjectiveCount, ConstraintCount, DecisionVariables int
	ConstraintTypes                                    []ConstraintType
	DecisionInterval                                   [][]float64
	Evaluate                                           func(Genotype) Fitness
	Name, TrueParetoFrontFilename                      string
	paretoFrontData                                    [][]float64
}

// TestSuite is a struct containing a set of CMOPs.
type TestSuite struct {
	NumberOfProblems int
	Problems         []Cmop
	Name             string
}

/*
TrueParetoFront returns the true paretofront in the form of a 2d array.
If no datafile for the pareto front is available, the function will return nil.
*/
func (c *Cmop) TrueParetoFront() [][]float64 {
	if c.paretoFrontData == nil && c.TrueParetoFrontFilename != "" {
		c.readParetoFrontFromFile()
	}
	return c.paretoFrontData
}

func (c *Cmop) readParetoFrontFromFile() {
	if data, err := filehandling.ParseDatFile(c.TrueParetoFrontFilename); err == nil {
		c.paretoFrontData = data
	} else {
		log.Println("Cmop:", "Error while parsing paretoFile")
	}
}
