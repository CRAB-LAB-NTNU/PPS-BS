# PPS-BS
Push Pull Search with Boundary Search

```
pps := pps.PPS{
		Cmop:                 cmop,
		Moea:                 &moead,
		Delta:                1.0 / 1000000.0,
		Epsilon:              1.0 / 1000.0,
		SearchingPreference:  0.95,
		ConstraintRelaxation: 0.1,
		RelaxationReduction:  2.0,
		TC:                   800,
		L:                    20,
	}
```

```
moead := optimisers.Moead{
		CMOP: cmop,
		WeightNeigbourhoodSize: 30,
		WeightDistribution:     300, // 23 for 3 objectives
		DecisionSize:           30,
		MaxChangeIndividuals:   2,
		GenerationMax:          300000,
		DEDifferentialWeight:   0.5,
		CrossoverRate:          1,
		DistributionIndex:      20.0,
	}
```

```
cmop := types.CMOP{
		NumberOfObjectives: 2,
		Calculate:          testSuite.CMOP1,
	}
```