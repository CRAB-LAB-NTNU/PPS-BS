package stages

import "github.com/CRAB-LAB-NTNU/PPS-BS/types"

type Binary struct {
	name      string
	isOver    bool
	stageType types.StageType
}

func NewBinary() *Binary {
	return &Binary{
		name:      "Binary",
		isOver:    false,
		stageType: types.BinarySearch,
	}
}

func (b Binary) Name() string {
	return b.name
}
func (b *Binary) Type() types.StageType {
	return b.stageType
}
func (b *Binary) SetOver() {
	b.isOver = true
}

func (b Binary) IsOver() bool {
	return b.isOver
}
