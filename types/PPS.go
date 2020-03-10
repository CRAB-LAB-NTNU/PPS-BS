package types

//Stage interface defines all the menthods needed to be a stage.
//Might be overkill to have this as an interface and not a single struct
type Stage interface {
	Stage() StageType
	SetOver()
	IsOver() bool
}

type StageType int

const (
	// Push PPS is in the push stage and constraints are ignored
	Push StageType = iota
	// BinarySearch PPS is a the stage between Push and Pull
	BinarySearch

	//Pull PPS is in the pull stage and constraints are handled
	Pull
)
