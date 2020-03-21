package plotter

import (
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
	"gonum.org/v1/plot/plotter"
)

func convertFloatToPoints(data []float64, switchPoint int) plotter.XYs {
	var points plotter.XYs
	for i, v := range data {
		if i < switchPoint {
			continue
		}
		point := plotter.XY{X: float64(i), Y: v}
		points = append(points, point)
	}
	return points
}

func convertParetoToPoints(pareto [][]float64, min, max float64) plotter.XYs {
	var points plotter.XYs
	for _, set := range pareto {
		if set[0] > max || set[1] > max || set[0] < min || set[1] < min {
			continue
		}
		point := plotter.XY{
			X: set[0],
			Y: set[1],
		}
		points = append(points, point)
	}
	return points
}

func convertIndividualsToPoints2D(individuals []types.Individual, min, max float64) plotter.XYs {
	var points plotter.XYs
	for _, ind := range individuals {
		objectiveValues := ind.Fitness().ObjectiveValues
		if objectiveValues[0] > max || objectiveValues[1] > max || objectiveValues[0] < min || objectiveValues[1] < min {
			continue
		}
		point := plotter.XY{
			X: objectiveValues[0],
			Y: objectiveValues[1],
		}
		points = append(points, point)
	}
	return points
}
