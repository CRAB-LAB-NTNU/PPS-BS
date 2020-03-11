package stages

import "github.com/CRAB-LAB-NTNU/PPS-BS/types"

type Pull struct {
	name   string
	isOver bool
	stage  types.StageType
}

func NewPull() *Pull {
	return &Pull{
		name:   "Pull",
		isOver: false,
		stage:  types.Pull,
	}
}

func (p Pull) Name() string {
	return p.name
}

func (p Pull) Stage() types.StageType {
	return p.stage
}

func (p *Pull) SetOver() {
	p.isOver = true
}

func (p Pull) IsOver() bool {
	return p.isOver
}
