package stages

import "github.com/CRAB-LAB-NTNU/PPS-BS/types"

type Pull struct {
	name      string
	isOver    bool
	stageType types.StageType
}

func NewPull() *Pull {
	return &Pull{
		name:      "Pull",
		isOver:    false,
		stageType: types.Pull,
	}
}

func (p Pull) Name() string {
	return p.name
}

func (p Pull) Type() types.StageType {
	return p.stageType
}

func (p *Pull) SetOver() {
	p.isOver = true
}

func (p Pull) IsOver() bool {
	return p.isOver
}
