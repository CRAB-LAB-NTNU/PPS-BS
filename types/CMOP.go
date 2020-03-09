package types

import (
	"log"

	"github.com/CRAB-LAB-NTNU/PPS-BS/filehandling"
)

// CMOP is a interface describing a multi objective optimisation problem
type CMOP struct {
	ObjectiveCount, ConstraintCount, DecisionVariables int
	ConstraintTypes                                    []ConstraintType
	DecisionInterval                                   [][]float64
	Evaluate                                           func(Genotype) Fitness
	Name, TrueParetoFrontFilename                      string
	paretoFrontData                                    [][]float64
}

/*
SetDecisionInterval sets the interval for all decision variables.
Should only be used if the interval is equal for all variables.
If the interval is not equal for all variables, set the attribute manually.
*/
func (c *CMOP) SetDecisionInterval(min, max float64) {
	c.DecisionInterval = make([][]float64, c.DecisionVariables)
	for i := range c.DecisionInterval {
		c.DecisionInterval[i] = append(c.DecisionInterval[i], min)
		c.DecisionInterval[i] = append(c.DecisionInterval[i], max)
	}
}

/*
TrueParetoFront returns the true paretofront in the form of a 2d array.
If no datafile for the pareto front is available, the function will return nil.
*/
func (c *CMOP) TrueParetoFront() [][]float64 {
	if c.paretoFrontData == nil && c.TrueParetoFrontFilename != "" {
		c.readParetoFrontFromFile()
	}
	return c.paretoFrontData
}

func (c *CMOP) readParetoFrontFromFile() {
	if data, err := filehandling.ParseDatFile(c.TrueParetoFrontFilename); err == nil {
		c.paretoFrontData = data
	} else {
		log.Println("CMOP:", "Error while parsing paretoFile")
	}
}
