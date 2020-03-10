package stages

import "github.com/CRAB-LAB-NTNU/PPS-BS/types"

type Binary struct {
	isOver bool
	stage  types.StageType
}

func (b *Binary) Stage() types.StageType {
	return b.stage
}
func (b *Binary) SetOver() {
	b.isOver = true
}

func (b Binary) IsOver() bool {
	return b.isOver
}
