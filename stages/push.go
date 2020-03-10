package stages

import (
	"math"

	"github.com/CRAB-LAB-NTNU/PPS-BS/arrays"

	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
)

type Push struct {
	rk, delta, epsilon float64
	l                  int
	isOver             bool
	stage              types.StageType
	ip, np             [][]float64
}

func NewPush(delta, epsilon float64, l int, generations int, decisionVariables int) *Push {
	return &Push{
		rk:      -1,
		delta:   delta,
		epsilon: epsilon,
		l:       l,
		isOver:  false,
		stage:   types.Push,
		ip:      arrays.Zeros2DFloat64(generations, decisionVariables),
		np:      arrays.Zeros2DFloat64(generations, decisionVariables),
	}
}

func (p Push) Stage() types.StageType {
	return p.stage
}
func (p *Push) SetOver() {
	p.isOver = true
}

func (p Push) IsOver() bool {
	return p.rk <= p.epsilon
}

func (p *Push) Update(generation int, population []types.Individual) {
	p.setIdealAndNadir(generation, population)
	p.calculateMaxChange(generation)
}

func (p *Push) setIdealAndNadir(generation int, population []types.Individual) {
	ip, np := biooperators.CalculateNadirAndIdealPoints(population)

	p.ip[generation] = ip
	p.np[generation] = np
}

func (p *Push) calculateMaxChange(generation int) {
	if generation >= p.l {
		rz := p.rx(generation, p.ip)
		rn := p.rx(generation, p.np)

		p.rk = math.Max(rz, rn)
	}
}

func (p Push) rx(g int, points [][]float64) float64 {
	m := math.SmallestNonzeroFloat64
	for i := 0; i < len(points[g]); i++ {
		cur := points[g][i]
		offset := points[g-p.l][i]
		dist := math.Abs(cur - offset)
		if calc := dist / math.Max(math.Abs(offset), p.delta); calc > m {
			m = calc
		}
	}
	return m
}
