package stages

import "github.com/CRAB-LAB-NTNU/PPS-BS/types"

type Binary struct {
	name   string
	isOver bool
	stage  types.StageType
}

func NewBinary() *Binary {
	return &Binary{
		name:   "Binary",
		isOver: false,
		stage:  types.BinarySearch,
	}
}

func (b Binary) Name() string {
	return b.name
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
