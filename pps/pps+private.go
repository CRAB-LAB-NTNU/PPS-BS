package pps

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/stages"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

func (pps *PPS) setIdealAndNadir(generation int) {
	ip, np := biooperators.CalculateNadirAndIdealPoints(pps.MOEA.Population())

	pps.idealPoints[generation] = ip
	pps.nadirPoints[generation] = np
}

func (pps *PPS) calculateMaxChange(generation int) {
	if generation >= pps.L {
		rz := pps.rx(generation, pps.idealPoints)
		rn := pps.rx(generation, pps.nadirPoints)
		pps.rk = math.Max(rz, rn)
	}
}

func (pps PPS) rx(k int, points [][]float64) float64 {
	m := math.SmallestNonzeroFloat64
	for i := 0; i < len(points[k]); i++ {
		cur := points[k][i]
		offset := points[k-pps.L][i]
		dist := math.Abs(cur - offset)
		if calc := dist / math.Max(math.Abs(offset), pps.Delta); calc > m {
			m = calc
		}
	}
	return m
}

func (pps *PPS) changeStage(generation int) bool {
	if pps.currentStage().Stage() == types.Push {
		return pps.changePush(generation)
	} else {
		return pps.currentStage().IsOver()
	}
	return false
}

func (pps *PPS) changePush(generation int) bool {
	push, ok := pps.currentStage().(*stages.Push)
	if !ok {
		panic("Could not assert push stage")
	}
	push.Update(generation, pps.MOEA.Population())
	pps.Stages[pps.stage] = push
	return pps.currentStage().IsOver()
}

func (pps PPS) currentStage() types.Stage {
	return pps.Stages[pps.stage]
}

func (pps *PPS) nextStage() {
	pps.stage++
}
