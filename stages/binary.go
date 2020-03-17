package stages

import "github.com/CRAB-LAB-NTNU/PPS-BS/types"

type Binary struct {
	name        string
	isOver      bool
	stageType   types.StageType
	minDistance float64
	fcp         float64
}

func NewBinary(minDistance, fcp float64) *Binary {
	return &Binary{
		name:        "Binary",
		isOver:      false,
		stageType:   types.BinarySearch,
		minDistance: minDistance,
		fcp:         fcp,
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

func (b Binary) MinDistance() float64 {
	return b.minDistance
}

func (b Binary) Fcp() float64 {
	return b.fcp
}
